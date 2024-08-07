---
title: "Service DB 에 분석 쿼리를 실행하면 어떤 문제가 발생할까?"
date: 2024-07-02T18:57:21+09:00
draft: false
categories :
- DataBase
- DataEngineering
---

Service DB는 사용자가 서비스를 사용하며 데이터를 읽거나 쓸 때 사용되는 DB이다. 서비스를 운영하다 보면 사용자의 데이터를 분석해야하는 상황이 발생한다.
예를 들어 온라인커머스 서비스를 운영하고 있다고 가정해보자. 그러면 온라인MD는 각 상품의 매출에 대한 지표를 확인하고 싶을 수 있다.

`sales` 라는 매출에 관련된 table이 있다면 온라인 MD 에게 상품의 매출에 대한 데이터를 제공해줄 때, 가장 간단한 방법은 `sales` table에 바로 쿼리를 작성하여 데이터를 제공해주는 것이다.
하지만 이때 발생하는 문제는 무엇일까?

오늘은 Service DB에 분석을 위한 쿼리를 실행했을 때 발생할 수 있는 문제와 해결할 수 있는 방법에 대해 생각해보고자 한다.

어떤 문제점이 발생할 수 있을지에 대해 알아보기 전, 먼저 어떤 지표를 확인하면 좋을지에 대해 생각해보았다.

### CPU
높은 CPU 사용률은 DB 서버에 많은 계산 작업을 수행하고 있다는 신호이다.

### Memory
메모리가 부족하다면 디스크 I/O가 증가하여 성능이 저할 될 수 있다.

### 디스크 I/O
쿼리가 많은 데이터를 읽거나 쓸 때, 디스크 I/O가 증가한다. 디스크 대기 시간이 길어지면 성능 저하가 발생할 수 있다.

### 쿼리 응답 시간
응답 시간이 길어지면, 성능 문제를 의심할 수 있다.

### 동시 연결 수
동시 연결 수가 많아지면 자원 경합이 발생할 수 있다.

### 락 대기 시간
쿼리 실행 중 락이 발생하여 다른 트랜잭션이 대기하는 시간을 모니터링해야 한다. 락 경합이 심하면 성능 저하가 발생할 수 있다.

현재 사내에서 사용하고 있는 DB는 PostgreSQL인데, 
PostgreSQL는 데이터를 읽을 때, 일반적으로 **공유 잠금 (shared lock)**이 설정되어 있기 때문에 여러 트랜잭션이 동시에 같은 행을 동시에 읽을 수 있다.
따라서, 락 대기 시간에 대한 지표는 확인하지 않아도 될 것 같다고 생각한다!

# 잠재적 문제
사용자가 더 증가하고 분석 요구사항이 더 많아진다면, 서비스 데이터의 양이 많아질 뿐만 아니라 쿼리 개수가 많아진다. 그러면 다음과 같은 문제가 발생할 수 있다.

- 서비스 사용자의 경험이 저하될 수 있다.
- CPU와 I/O 리소스를 더 많이 사용하여 데이터베이스 전체 성능을 저하시킬 수 있다.
- 더 심각한 부하가 발생한다면 데이터베이스가 다운될 수 있다.
- 데이터베이스 부하로 인해 트랜잭션이 실패할 수 있으며, 심각한 경우 데이터 손실이 발생할 수 있다.

따라서 실제 서비스에 사용되는 DB와 분석을 위한 DB를 분리하는 것이 필요하다.

# 목표
다음과 같은 목표를 수립할 수 있다.

- 서비스 DB는 서비스를 위해 운영한다. 
  - 분석은 서비스 DB에 접근하지 않는다.
- 데이터 무결성이 깨지지 않아야 한다.
- Optional: 금액이 크게 증가하지 않는다.

# 해결 방법
## 1. 데이터 복제 (with RDBMS)
서비스 DB의 데이터를 분석 DB로 복제하여 두 개의 DB를 관리하는 방법이다.

### 1-1. Pub/Sub & Subscribe Server
Google Cloud Pub/Sub과 같은 메시징 시스템을 사용하여 데이터 변경 이벤트를 분석 DB로 스트리밍하는 방법이다.
![image](https://github.com/yumin00/blog/assets/130362583/13f3dc5d-e2ef-4153-967b-db92abb3af17)

Data 변경 이벤트가 발생하면, API 서버는 비동기적으로 Google Pub/Sub에 변경된 데이터를 Message로 Publish하고, 또 다른 Server가 해당 Message를 Subscribe하여 Analytics DB에 데이터를 업데이트하는 방법이다.

[장점]

- API 서버에서 직접 데이터 변경 이벤트를 감지하므로 구현이 비교적 간단하다.
- 응답이 실패하면 Message를 Publish하지 않기 때문에 데이터 복제의 신뢰성이 높다.

[단점]
- API 서버에서 변경된 데이터를 감지하여 Message 를 Publish 해야 하기 때문에 시스템 복잡성이 증가한다.
- Subscribe하여 Analytics DB에 데이터를 복제하는 Server를 추가 구현해야 한다.
- 데이터 변경 이벤트와 Message Publish 를 모두 API 서버에서 관리해야 때문에, 시스템 결합도가 높아져 유지보수가 어려워질 수 있다.
- Service DB의 스키마 버전, Google Pub/Sub 스키마 버전, Bigqeury 스키마 버전을 관리하지 않으면 스키마 불일치로 인한 오류가 발생할 수 있다.

### 1-2. Airflow
Airflow를 사용하여 서비스 DB의 데이터 변경을 감지하고 이를 분석 DB에 복제하는 방법이다.

![image](https://github.com/yumin00/blog/assets/130362583/67c1c2e6-13f7-4a2d-8ad2-9ab2d31e89b2)

Service DB의 데이터 변경 이벤트를 Trigger하는 DAG를 생성하고, 데이터 변경 이벤트가 발생하면 이를 분석 DB에 복제하는 방법이다.

[장점]
- 데이터 변경 이벤트와 API 서버를 분리하여 관리할 수 있다.
- Airflow에서 DB 데이터 감지 기능을 제공하므로 별도의 변경 감지 시스템을 구현할 필요가 없다.
- Airflow의 모니터링 기능을 통해 데이터변경이 실시간으로 처리되는지 쉽게 확인할 수 있다.

[단점]
- Airlfow를 사용하고 있지 않다면 초기 설정 및 구성에 대한 러닝커브가 존재한다.
- Airflow 운영을 위해서 추가적인 인프라 및 리소스가 필요할 수 있다.

## 2. 데이터 복제 (with BigQuery)
데이터를 복제하되 1번과 다르게 PostgreSQL이 아닌 BigQuery를 사용하여 복제하는 방법이다.

### 2-1. Google Pub/Sub (ETL)
Google Cloud Pub/Sub과 같은 메시지 패싱 시스템을 사용하여 데이터 변경 이벤트를 전송하고, 변경된 데이터를 BigQuery 로 스트리밍하는 방법이다.
![image](https://github.com/yumin00/blog/assets/130362583/a8a18e8c-6cc6-4fb2-b221-a123475d2ceb)

[장점]
- API 서버에서 직접 데이터 변경 이벤트를 감지하므로 구현이 비교적 간단하다.
- 응답이 실패하면 Message를 Publish하지 않기 때문에 데이터 복제의 신뢰성이 높다.
- 분석 DB에 데이터를 적재하는 서버를 별도로 구축하지 않고, Google Function을 통해 데이터를 적재할 수 있다.
- 대규모 데이터 셋일 경우, RDBMS보다 더 빠른 쿼리 속도를 얻을 수 있다.

[단점]
- API 서버에서 변경된 데이터를 감지하여 Message 를 Publish 해야 하기 때문에 시스템 복잡성이 증가한다.
- 데이터 변경 이벤트와 Message Publish 를 모두 API 서버에서 관리해야 때문에, 시스템 결합도가 높아져 유지보수가 어려워질 수 있다.
- Service DB의 스키마 버전, Google Pub/Sub 스키마 버전, Bigqeury 스키마 버전을 관리하지 않으면 스키마 불일치로 인한 오류가 발생할 수 있다.
- Google Function의 버전 관리를 필요로 한다.

## 3. ETL 프로세스
00시(예시)에 Service DB에서 필요한 데이터를 추출하여 분석 요구사항에 맞춰 변환한 후 이를 분석 DB에 저장하는 방식이다.
![image](https://github.com/yumin00/blog/assets/130362583/ae65cca8-427f-46ff-8286-0228dca2693f)

[장점]
- 데이터 변경 이벤트와 API 서버가 결합되어 있지 않아 강결합 문제가 발생하지 않는다.
- 데이터 분석 목적에 맞게 변환할 수 있어, 분석 DB의 구조를 최적화할 수 있다.

[단점]
- 주로 배치 작업으로 수행되기 때문에, 실시간 분석이 어려우며 데이터의 최신성을 유지할 수 없다.
- 요구사항이 발생할 때마다 Job과 분석 DB를 수정해야 한다.

## 4. AuditLog 사용
기존 Pub/bSub에 쌓이던 AuditLog를 사용하는 프로세스입니다.
![image](https://github.com/yumin00/blog/assets/130362583/e34506d3-cce7-4e0a-ae64-7e0cff8d6ba9)

[장점]

- 기존 API 서버에서 AuditLog를 Pub/Sub 으로 Publish 하고 있기 때문에 API 서버의 리소스가 필요하지 않습니다.
- Bigquery에서 요구사항에 맞춰 테이블을 따로 설정할 경우, 쿼리의 지연 속도 시간을 최소화할 수 있습니다.

[단점]

- AuditLog를 원천 데이터로 사용하기 때문에, API 서버에 문제가 발생할 경우 분석 데이터도 영향을 받습니다.
- Google Functions의 추가 구현이 필요하며, 버전 관리도 필요합니다.

## 5. CDC (Change Data Capture)
### 5-1. Google DataStream
Google Cloud에서 제공하는 서비스인 Data Stream을 사용하는 방식입니다. DataStream은 Change Data Capture (CDC) 기술을 사용하여 실시간으로 데이터베이스의 변경 사항을 추적하고,
이를 대상으로 데이터를 다른 시스템으로 전송하는 기능을 제공합니다.Data Stream을 통해 PostgreSQL 데이터베이스의 변경 사항을 감지하고 이를 Google BigQuery로 복제할 수 있습니다.

![image](https://github.com/user-attachments/assets/745b7243-6739-4669-8331-7f84bc21aaf6)

[장점]
- Google Cloud Console에서 추가 코드 작성 없이 설정을 통해 진행할 수 있습니다.
- 실시간으로 데이터 변경 사항을 추적하기 때문에, 데이터 유실의 위험성이 적습니다.
- Service DB의 데이터 변경 사항을 추적하기 때문에, 데이터 무결성이 깨지지 않습니다.
- Google Cloud의 관리형 서비스로, 인프라 관리 부담이 적습니다.

[단점]
- Service DB의 데이터를 감지하여 Bigquery에 적재하는 방식이기 때문에 커스터마이징이 어렵습니다.
  - 즉, Service DB의 테이블 형태 그대로 저장해야 합니다.

### 5-2. Google Pub/Sub & Google Functions 
Google Cloud Pub/Sub와 Google Functions를 사용하여 PostgreSQL의 데이터 변경 사항을 BigQuery로 스트리밍하는 방법입니다.
PostgreSQL의 변경 사항을 Pub/Sub으로 메세지를 게시하여, 해당 메세지를 Google Functions를 통해 BigQuery에 적재하는 방식입니다.

![image](https://github.com/user-attachments/assets/3635b73d-9dfb-45b4-a88c-8bfb1e482f0c)

[장점]
- 트리거와 함수 코드를 커스터마이징하여 특정 요구사항에 맞게 조정할 수 있습니다.
- Service DB의 데이터 변경 사항을 추적하기 때문에, 데이터 무결성이 깨지지 않습니다.
- Google Cloud의 관리형 서비스로, 인프라 관리 부담이 적습니다.

[단점]
- PostgreSQL에서 직접 트리거를 설정해야 합니다.
- Google Functions의 추가 코드 작성이 필요합니다.
- Google Pub/Sub 스키마 버전을 관리하지 않으면 스키마 불일치로 인한 오류가 발생할 수 있습니다.
- Google Functions의 버전 관리를 위하여 배포 프로세스가 필요합니다.

# 선정한 아키텍처
분석 DB의 목적은 서비스 DB의 안전성을 해치지 않고 데이터 분석을 위한 것이기 때문에, 서비스 DB와 분리하지만 서비스 DB의 데이터를 원천으로 데이터 무결성이 깨지지 않아야 합니다.
또한 해당 작업에 너무 많은 리소스가 투입되지 않기 위해서 선정한 아키텍처는 5-1 DataStream 입니다.

![image](https://github.com/user-attachments/assets/745b7243-6739-4669-8331-7f84bc21aaf6)

서비스 DB의 데이터 변경 감지를 통해  BigQuery에 저장하는 방식이기 때문에 데이터 무결성을 지킬 수 있고, RDBMS가 아닌 BigQuery를 사용하기 때문에 금액적으로 큰 부담이 없습니다.