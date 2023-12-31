---
title: "Kafka Producer에 대해 알아보자"
date: 2023-10-18T21:43:01+09:00
draft: false
---

# Kafka Producer
## 1. 들어가며
회사에서 기존에 메시지를 담기 위해 RabbitMQ를 사용하고 있었다. RabbitMQ는 메시지 브로커로 프로듀서가 메시지를 넣고, 컨수머가 메시지를 가져갈 수 있게 해주는 것이다.

RabbitMQ와 Kafka의 차이점은 메시지의 영속성 보장 여부이다. RabbitMQ는 컨슈머가 메시지를 가져가면 소실되고, Kafka는 메시지를 파일 시스템에 저장함으로써 영속성을 보장한다.

회사에서 RabbitMQ에서 Kafka로 이전하는 작업을 진행하게 되어, 먼저 Kafka에 대해 공부해보고자 한다.

## 2. Kafka 구성
Kafka는 메시지를 생산, 발송하는 프로듀서(producer)와 메시지를 소비, 수신하는 컨슈머(consumer), 그리고 프로듀서와 컨슈머 사이에 메시지를 중개하는 브로커(Broker)로 구성된다.

이번에는 Kafka의 Producer에 대해 알아보고자 한다.

## 3. Kafka Producer
### kafka 프로듀서란?
프로듀서는 보통 Kafka 프로듀서 API와 그것으로 구성된 애플리케이션을 말한다. 그리고 프로듀서는 브로커에 특정 토픽(혹은 파티션 영역까지)을 지정하여 메시지를 전달하는 역할을 담당한디.

### 프로듀서를 통해 전달되는 메시지의 구조
- 토픽 (Topic)
- 토픽 중 특정 파티션 위치 (Partition)
- 메시지 생성 시간 (Timestamp)
- 메시지 키 (Key)
- 메시지 값 (Value)

### 메시지 전달 과정
<img width="742" alt="image" src="https://github.com/yumin00/blog/assets/130362583/ef87a927-e464-45fb-a9b9-ed3f3d1d1797">

출처 : https://dzone.com/articles/take-a-deep-dive-into-kafka-producer-api

프로듀서틑 4가지 과정을 통해 메시지를 브로커에게 전달한다.

1. 직렬화 (Serializer)

프로듀서는 전달 요청 받은 메시지를 직렬화한다. 직렬화는 Serializer가 지정된 설정을 통해 처리하며, 메시지의 키와 값은 바이트 뭉치 형태로 변환된다.

> 직렬화란?
> 
> 객체의 상태를 보관이나 전송 가능한 형태로 변환하는 것
> 
> - `json` 라이브러리의 `Marshal()` 함수로 JSON 형태로 직렬화

2. 파티셔닝 (Partitioner)

직렬화된 메시지는 Partitioner를 통해 정의된 로직에 따라 토픽의 어떤 파티션에 저장될지 결정된다.

별도의 Partitioner 설정을 하지 않으면 Round Robbin 형태로 파티셔닝을 진행한다. 즉, 파티션들에게 골고루 전달할 수 있도록 파티셔닝을 한다.

3. 압축 (Compression)

만약 메시지 압축이 설정되어 있다면, 설정된 포맷에 맞춰 메시지를 압축한다.

압축된 메시지는 브로커에게 빠르게 전달할 수 있을 뿐만 아니라 브로커 내부에서도 빠르게 복제가 가능하다.

4. 전달 (Sender)
  
프로듀서는 메시지를 TCP를 통해 브로커 리더 파티션으로 전송한다. 메시지마다 매번 네트워크를 통해 전달하는 것은 비효율적이기 때문에 지정된만큼 메시지를 저장했다가 한번에 브로커에게 전달한다.

이 과정은 프로듀서 내부의 Record Accumulator(RA)가 담당하여 처리한다.

RA는 각 토픽 파티션에 대응하여 배치 큐(Batch Queue)를 구성하고, 메시지들을 레코드 배치(Record Batch) 형태로 묶어 큐에 전달한다.

각 배치 큐에 저장된 레코드 배치들은 때가 되면 각각 브로커에게 전달된다. 이 과정은 Sender가 처리한다.

Sender는 스레드 형태로 구성되며, 관리자가 설정한 특정 조건에 만족한 레코드 배치를 브로커로 전송한다.

이때, 네트워크 비용을 줄이기 위해 piggyback 방식으로 조건을 만족하지 않은 다른 레코드 배치를 조건을 만족한 것과 함께 브로커를 전송한다.

예를 들어, 3번 브로커로 전송되어야 하는 `가 토픽`의 `파티션 B`의 큐에 레코드 배치가 전송할 조건을 만족했을 경우 Sender는 해당 레코드 배치를 가져와 3번 브로커로 전송할 준비를 한다.
이때, 3번 브로커로 전송되어야 하는 `다 토픽`의 `파티션 A`의 큐의 레코드 배치가 전송할 조건을 만족하지 않았더라도 Sender는 이 레코드 배치를 업어 함께 3번 브로커로 전송한다.

5. 응답

브로커에 네트워크 전송 요청을 보낸 Sender는 설정 값에 따라 브로커의 응답을 기다리거나 가다리지 않는다.

- 브로커의 응답을 기다리지 않는 설정일 경우, 메시지 전송에 대한 과정이 마쳐진다.
- 브로커의 응답을 기다리는 설정일 경우, 메시지 전송 성공 여부를 응답으로 받는다.
  - 실패 응답을 받을 경우, 설정 값에 따라 재시도를 시도한다.
    - 재시도 횟수를 초과한 경우, 예외를 뱉어낸다.
  - 성공 응답을 받을 경우, 메시지가 저장된 정보(메타데이터)를 반환한다.
    - 메타데이터는 메시지가 저장된 토픽, 파티션, 오프셋, 타임스탬프 정보를 가지고 있다.

## 4. Kafka Producer 옵션
- key.serializer 및 value.serializer

프로듀서가 키와 값을 직렬화하는 데 사용하는 클래스를 정의한다. 이 클래스들은 org.apache.kafka.common.serialization.Serializer 인터페이스를 구현해야 한다.

- bootstrap.servers

- Kafka 클러스터에 처음 연결하기 위해 사용되는 호스트/포트 쌍의 목록입니다. 클라이언트는 이 서버들을 사용하여 전체 서버 세트를 발견한다.

- buffer.memory

프로듀서가 서버로 전송 대기 중인 레코드를 버퍼링하는 데 사용할 수 있는 메모리의 총 바이트 수

- compression.type

프로듀서가 생성하는 모든 데이터의 압축 유형이다. 유효한 값은 'none', 'gzip', 'snappy', 'lz4', 'zstd'이다.

- retries

잠재적으로 일시적인 오류로 전송이 실패한 레코드를 재전송할 횟수

- batch.size

프로듀서가 같은 파티션으로 전송하는 여러 레코드를 하나의 요청으로 묶으려고 시도하는 기본 배치 크기를 제어

- linger.ms

프로듀서가 단일 배치 요청으로 그룹화하기 위해 레코드 전송 사이에 대기하는 시간

- client.id

서버에 요청을 할 때 전달되는 ID 문자열이다.. 이는 요청의 출처를 추적하는 데 사용됩니다.

- max.request.size

바이트 단위의 요청의 최대 크기이다. 이 설정은 프로듀서가 단일 요청으로 보낼 수 있는 레코드 배치의 수를 제한한다.

- ssl.key.password, ssl.keystore.location 등

SSL 구성에 사용되는 여러 파라미터들이 있으며, 이들은 클러스터와의 보안 연결을 설정하는 데 사용된다.

## 5. Kafka Transaction Producer
### Kafka Transaction Producer
트랜잭션 프로듀서는 모든 데이터의 원자성을 만족시키기 위한 옵션이다. 다수의 데이터를 트랜잭션으로 묶음으로써 전체 데이터를 처리하거나 처리하지 않도록 만드는 방식이다.

트랜잭션 프로듀서는 컨슈머로 하여금 데이터를 그냥 가져가는 것이 아니라 트랜잭션 처리가 완료된, 즉 commit 상태인 데이터만 가져가게 한다.

트랜잭션 프로듀서는 파티션에 레코드가 저장될 때 commit레코드를 보내주어 해당 데이터를 가져가도 된다는 것을 표시한다.

## 6. Kafka Idempotence Producer
멱등성 프로듀서는 여러 번 연산을 수행해도 동일한 결과를 나타내는 것을 의미한다.

멱등성 프로듀서는 데이터를 여러 번 전송해도 브로커에는 딱 한 개의 데이터 저장됨을 의미한다.

enable.idempotence 옵션을 true로 하여 지정할 수 있다.

데이터를 브로커롤 전달할 때 프로듀서의 PID와 시퀀스 넘버를 함께 전달하는데, 그 결과 브로커는 PID와 시퀀스 넘버를 확인하여 동일한 메시지의 적재 요처이 와도 딱 한 번만 데이터를 적재하도록 한다.