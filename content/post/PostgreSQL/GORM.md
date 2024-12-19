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

1. Cloud Run 컨테이너 시작 시:
- Cloud Run은 컨테이너가 시작되기 전에 `/cloudsql` 디렉터리를 컨테이너 파일 시스템에 마운트
- Cloud SQL Auth Proxy 초기화

2. Unix 소켓 파일 생성
- Cloud SQL Auth Proxy는 `/cloudsql/<INSTANCE_CONNECTION_NAME>` 경로에 Unix 도메인 소켓 파일 생성
  - 이 소켓 파일은 실제 데이터가 저장된 파일이 아니라 Cloud SQL Auth Proxy와 통신하는 인터페이스
- Cloud Run 서비스 계정의 IAM 권한을 확인하여 Cloud SQL에 연결할 수 있는 인증 설정
- Proxy는 Google Cloud 내부 네트워크를 통해 지정된 Cloud SQL 인스턴스와 연결을 설정합니다.

3. 데이터 요청:
- Unix 소켓 파일에 데이터를 쓰면, Proxy가 이 데이터를 수신
- Proxy는 데이터베이스 요청을 Google Cloud 네트워크를 통해 지정된 Cloud SQL 인스턴스로 전송

4. Cloud SQL 처리:
- Proxy를 통해 전달된 요청은 Cloud SQL 인스턴스에서 처리
- 처리된 결과는 Google Cloud 네트워크를 통해 다시 Proxy로 전달됩니다.

4. 응답 수신:
- 소켓 파일을 통해 응답 데이터를 받아 Cloud Run이 이를 처리한다.
- Google Cloud는 /cloudsql 디렉터리를 컨테이너 내부에 마운트한다.
- 사용자가 설정한 Cloud SQL 인스턴스(INSTANCE_CONNECTION_NAME)와 연결되는 Unix 도메인 소켓 파일을 `/cloudsql/<INSTANCE_CONNECTION_NAME>` 경로에 생성한다.
- 컨테이너 내부에서 이 소켓 파일을 마치 로컬 파일처럼 접근하여 Cloud SQL과 통신할 수 있다.
- 이 파일은 Cloud Run과 Google Cloud 관리형 네트워크를 연결하는 가상 통로 역할을 한다.
- 로컬 파일처럼 보이지만, 이 소켓 파일은 실제 데이터 저장소가 아니며 Google Cloud 내부 네트워크와의 가상 연결이다.


### 장점
- 네트워크를 타지 않기 때문에 데이터 전송이 더 빠르고 네트워크 오버헤드가 줄어든다는 장점이 있다.
- Google Cloud의 내부 관리 네트워크를 통해 이루어지며 외부 네트워크(인터넷)를 사용하지 않아 데이터 보호 수준을 높일 수 있다.