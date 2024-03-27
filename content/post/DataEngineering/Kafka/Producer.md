---
title: "Kafka의 Producer에 대해 알아보자"
date: 2024-03-13T22:23:10+09:00
draft: false
categories :
- DataEngineering
- Kafka
---

패스트캠퍼의 `한번에 끝내는 Kafka Ecosystem` 강의 내용을 바탕으로 Producer를 정리해보고자 한다.

# Producer
Producer는 메시지를 생산해서 카프카 토픽으로 메시지를 보내는 애플리케이션을 말한다.

## Message 구조
message의 구조는 다음과 같다. (message == recore == event == data)

- headers : 메타데이터(topic, partition, timestamp, etc) 
- key : body
- value : body

이때, key, value는 avro, json 등 다양한 형태가 가능하며, 카프카에서는 메시지를 byte array 로 저장한다.

Producer가 json, string, avor, protobf 등 다양한 형태로 메시지를 프로듀스할 수 있는데, 이때 **Serializers** 를 지정해야 한다.

### Serializers
![image](https://github.com/yumin00/blog/assets/130362583/f3bb5022-ec3b-412b-a5ae-7368f32be4f9)

Producer가 메시지를 프로듀스할 때, serializer 를 선택하여 메시지를 byte array로 변형한 뒤 카프카에 저장한다.

- StringSerializer: 문자열 데이터를 byte array로 직렬화
- IntegerSerializer: 정수 데이터를 byte array로 직렬화
- ByteArraySerializer: byte array 데이터를 그대로 사용, 별도의 직렬화 과정이 필요하지 않을 때 사용

BOOTSTRAP_SERVERS를 지정하는 것처럼 KEY_SERIALIZER_CLASS_CONFIG / VALUE_SERIALIZER_CLASS_CONFIG 를 지정할 수 있는데, key 와 value의 serializers를 별도로 지정할 수 있다.

Seralizers를 통해 데이터를 직렬화함으로써 consumer가 데이터를 올바르게 처리할 수 있도록 하며, 네트워크를 통해 데이터를 전송할 때 데이터 크기를 최소화할 수 있다.

### Deserializers
Consumer가 컨슘할 때는 deserializers를 사용하여 데이터를 뽑을 수 있다.

## Produce 동작 원리
![image](https://github.com/yumin00/blog/assets/130362583/e24cb50b-b799-49dd-a730-38e5f91a4c77)

1. 메시지를 produce하기 위해서는 Producer Record를 생성해야 한다.
Producer Record : Topic / Partition / Timestamp / Key / Value / Header : topic, value는 필수

2. 메시지 send()

3. 지정한 Serializer 을 통해 byte array로 변경된다.

4. Partitioner 를 통해서 어떤 파티션으로 갈지 결정된다.

5. compress option: 압축 옵션을 썼으면 압축 여부 결정된다.

6. RecordAccumulator 에서 데이터를 묶음 형태로 모아진다.

7. 옵션에 따라서, 단건 or 설정한 개수 / 시간에 맞춰서 카프카로 전송된다.

8. 카프카가 받아서 ack를 보내준다. (실패 or 성공)

9. 실패하면 재시도 옵션이 있다.

10. 성공하면 메타데이터를 프로듀서에게 전달한다.

## Partitioner 역할
메시지를 전송 받으면, 해당 메시지를 토픽 내의 파티션 중에 어떤 파티션에 보낼지 결정하는 역할을 한다.(Partitioner는 토픽 내에 몇 개의 파티션이 있는지 알고 있다)

Partitioner의 전재 조건은 key가 null이 아닐 경우이다. key가 null이라면 Partitioner는 작동하지 않는다. 즉, key에 따라서 파티션이 나뉘어서 들어가지며 같은 key를 가진 메시지들은 같은 파티션에 적재된다고 할 수 있다.

```partition = hash(key) % number of parititons```

그렇다면 key가 null일 경우에는 어떻게 파티션이 결정될까? 카프카 2.4 이전 버전의 경우에는 라운드 로빈으로 들어갔지만, 2.4 이후에는 sticky 정책을 적용하였다.

sticky 정책이란, 한 파티션의 한 배치가 닫힐 때까지 한 파티션에만 보내고, 해당 파티션의 배치가 다 차면 랜덤으로 다른 파티션에 보내는 것이다.

Partitioner는 개발을 통해 직접 customizing이 가능하다.