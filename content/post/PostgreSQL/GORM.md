---
title: "GORM을 사용하여 PostgreSQL에 연결하기"
date: 2024-12-18T19:41:14+09:00
draft: true
categories :
- PosrgreSQL
---

## Cloud Run과 Cloud SQL
Cloud Run에서 Cloud SQl을 연결하는 방법에는 두 가지가 있다.
- Unix Socket으로 연결
- Cloud SQL Connector와 연

## Unix Socket
### Socket (소켓)
소켓이란 프로세스가 서로 연결하고 통신하기 위한 엔드포인트이다. 즉 소켓을 열어서 다른 프로세스와 통신할 수 있도록 개방하는 것이다.
소켓을 열면, 해당 소켓으로 특정 주소와 포트를 사용해서 접근할 수 있다.

### Unix Socket (유닉스 소켓)
유닉스 소켓은 IPC(Inter Process Communication) 방식이라고도 한다. 동일한 컴퓨터에서 실행되는 프로세스 간의 양방향 통신을 위한 방법이다.

유닉스 소켓 방식은 네트워크를 사용하여 통신하는 것이 아니라, 운영 체제의 커널에 있는 프로세스가 서로 통신할 수 있다.

즉, TCP/IP와 같은 방식을 사용하려면 서로의 주소를 알고 있어야하지만, 유닉스 소켓을 사용하면 주소 없이도 파일 시스템 내에서 통신이 가능하다.

## Unix Socket을 통한 Cloud Run과 Cloud SQL의 통신 방법
유닉스 소켓에 대해 파악했으니, Cloud Run과 Cloud SQL이 통신할 때 유닉스 소켓을 사용하면 어떤 방식을 통해 통신하게 되는지 알아보자.

Cloud Run에서 Cloud SQL과 연결하기 위해 `/cloudsql/<INSTANCE_CONNECTION_NAME>` 과 같은 이름을 사용하여 유닉스 도메인 소켓 파일을 생성한다.

유닉스 도메인 소켓 파일이란 Cloud SQL과 연결된 가상 소켓 파일로 Google Cloud의 관리형 네트워크와 연결되어 Cloud SQL과 통신할 수 있게 한다. 이 파일은 Cloud Run과 Cloud SQL 간의 데이터 전송 경로로 사용된다.

Google Cloud의 내부 네트워크는 소켓 파일을 통해 받은 데이터를 Cloud SQL로 전달하고, 반대로 Cloud SQL에서 반환된 데이터를 소켓 파일을 통해 Cloud Run에 전달한다.

정리해보자면,

![image](https://github.com/user-attachments/assets/c0a47d37-6cfc-4263-adf8-2eaa47f38ab1)

1. Cloud Run 컨테이너 시작 시:
- Cloud Run 인스턴스가 시작되면 /cloudsql/<INSTANCE_CONNECTION_NAME> 경로에 Unix 소켓 파일이 생성됨
  - 소켓은 Proxy Client에 해당되며, 컨테이너 내부에서 애플리케이션과 Proxy가 통신하는 인터페이스이다.

2. Proxy Client 초기화
- Cloud SQL Auth Proxy가 컨테이너 내부에서 실행되며, 해당 소켓 파일을 통해 애플리케이션 요청을 수신할 준비를 함
- Proxy Client는 Cloud Run 서비스 계정을 사용하여 Google Cloud에 인증을 진행한다.

3. Proxy Client와 Proxy Server 연결:
- Proxy Client는 Google Cloud 네트워크를 통해 Cloud SQL의 Proxy Server와 보안 터널(TCP Secure Tunnel)을 설정한다.
- 그림과 같이 TLS 암호화를 사용하여 안전하게 데이터를 주고받는다.

4. 데이터 요청:
- Unix 소켓 파일에 데이터를 쓰면, Client Proxy가 이 데이터를 수신
- Client Proxy는 데이터베이스 요청을 Google Cloud 네트워크를 통해 지정된 Cloud SQL Proxy Server로 전송
- Proxy Server는 Proxy Client로부터 수신한 요청을 Cloud SQL 데이터베이스로 전달합니다.

5. Cloud SQL 처리:
- Proxy를 통해 전달된 요청은 Cloud SQL 인스턴스에서 처리
- 처리된 결과를 Proxy Server에 전달하고, Google Cloud 네트워크를 통해 다시 Client Proxy로 전달된다.

6. 응답 수신:
- 소켓 파일을 통해 응답 데이터를 받아 Cloud Run이 이를 처리한다.


### 장점
- 네트워크를 타지 않기 때문에 데이터 전송이 더 빠르고 네트워크 오버헤드가 줄어든다는 장점이 있다.
- TCP 주소나 포트 번호를 따로 관리하지 않아도 되며 단순히 Unix 파일 시스템 경로를 설정하면 되기 때문에 복잡성을 줄이고 설정 작업을 간소화할 수 있다.
-  Unix 소켓 파일에 대한 접근은 컨테이너 내부의 파일 시스템 권한으로 제어되므로 불필요한 보안 위협이 줄어든다.

### 단점
- Cloud SQL의 연결 수에 제한이 있기 때문에 Cloud Run 인스턴스가 많이 스케일링 되면 연결 수가 초과될 수 있다.
  - 연결 풀(pool)과 같은 관리가 제대로 이루어지지 않으면 연결이 끊기거나 연결 실패가 발생할 수 있다.

## Cloud SQL Connector
![image](https://github.com/user-attachments/assets/5b96dc55-90ba-41ab-b6d8-8ad096ab0f4b)


1. 애플리케이션 초기화:
- Cloud Run 컨테이너가 시작되면 애플리케이션 코드에서 Cloud SQL Connector 라이브러리를 초기화
- 라이브러리는 Cloud Run의 서비스 계정 credentials을 자동으로 로드

2. 커넥션 풀 생성:
- Connector 라이브러리가 데이터베이스 커넥션 풀을 생성
- 커넥션 풀은 효율적인 데이터베이스 연결 관리를 위해 연결을 재사용

3. IAM 인증 처리:
- 데이터베이스 연결 요청 시 Connector가 자동으로 IAM 인증을 처리
- OAuth 2.0 access token을 생성하여 데이터베이스 연결 인증에 사용
- 토큰은 자동으로 refresh되어 지속적인 연결 유지

4. 보안 터널 설정:
- Connector는 내부적으로 Google Cloud의 IAP(Identity-Aware Proxy) 터널을 사용
- TLS 암호화된 연결을 통해 Cloud SQL 인스턴스와 직접 통신
- 이 과정에서 별도의 Proxy 서버가 필요하지 않음

5. 데이터베이스 요청 처리:
- 애플리케이션의 쿼리 요청이 커넥터를 통해 직접 Cloud SQL로 전달
- Cloud SQL이 요청을 처리하고 결과를 같은 보안 터널을 통해 반환

6. 연결 관리:
- Connector가 자동으로 연결 상태를 모니터링
- 연결 문제 발생 시 자동 재연결 시도
- 커넥션 풀의 크기를 자동으로 조절하여 최적의 성능 유지

Unix Socket 방식과 비교했을 때 주요 차이점은 다음과 같다.

- 별도의 Proxy 프로세스가 필요 없음
- IAM 인증이 라이브러리 레벨에서 자동으로 처리
- Cloud Run의 오토스케일링 환경에서 더 안정적으로 동작

이런 특징들 때문에 새로운 애플리케이션 개발 시에는 Cloud SQL Connector 방식이 더 권장된다.

Cloud SQL Connector를 사용하면 Cloud Run의 오토스케일링 환경에서 더 안정적으로 동작할 수 있는데 그 이유는 다음과 같다.

Unix Socket 방식의 경우:
- 새로운 Cloud Run 인스턴스가 생성될 때마다:
  - Auth Proxy 컨테이너가 시작되어야 함 
  - Unix 소켓 파일이 생성되어야 함 
  - Proxy가 초기화되고 인증을 완료해야 함 
  - 이 모든 과정이 완료된 후에야 데이터베이스 연결이 가능

Cloud SQL Connector 방식의 경우:
- 새로운 Cloud Run 인스턴스가 생성될 때:
  - 라이브러리가 초기화되면서 바로 연결 준비 
  - 별도의 프로세스 시작이나 소켓 파일 생성 불필요 
  - IAM 인증이 비동기적으로 처리됨



이러한 차이로 인해 다음과 같은 이점이 잇다.:

- 인스턴스 시작 시간이 더 빠름
- 메모리 사용량이 더 적음 (별도 Proxy 프로세스가 없으므로)
- 인스턴스 스케일 아웃 시 연결 설정이 더 빠르고 안정적
- 컨테이너 장애 발생 시 복구가 더 간단함 (단일 프로세스이므로)

특히 트래픽이 급증하는 상황에서 빠른 스케일 아웃이 필요할 때, 이러한 차이가 성능과 안정성에 영향을 미칠 수 있습니다.
