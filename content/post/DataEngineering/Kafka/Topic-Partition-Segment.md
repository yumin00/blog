---
title: "Kafka Topic에 대해 알아보자"
date: 2024-03-12T22:39:24+09:00
draft: true
categories :
- DataEngineering
- Kafka
---

패스트캠퍼의 `한번에 끝내는 Kafka Ecosystem` 강의를 들은 내용을 바탕으로 공부한 내용을 요약 정리해보고자 한다.

# 1. Kafka 주요 요소
Topic에 대해 알아보기 전, 카프카의 주요 요소에 대해 먼저 간단히 알아보자.

![image](https://github.com/yumin00/blog/assets/130362583/67cc48ca-556b-4266-af65-063ae45e5c68)

## Topic
Kafka 안에서 메시지가 저장되는 장소

## Producer
메지를 생산해서 Topic 으로 메시지를 보내는 애플리케이션

## Consumer
- 메시지를 소비하는 애플리케이션
- 하나의 consumer는 하나의 group 안에 들어감

## Consumer Group
- topic의 메시지를 사용하기 위해 협력하는 consumer들의 집합
- 해당 group의 consumer들은 하나의 세트처럼 움직이게 됨
- group 내의 consumer들은 협력하여 topic 메시지를 분산 병렬 처리함

> 분산 병렬 처리
> 
> 각 consumer들이 동시에 메시지를 consume하는 방식

## Producer와 Consumer의 분리 (decoupling)
producer와 consumer는 서로 완전히 분리되어 있기 때문에 누가 produce 하는지 누가 consume 하는지 서로 알 수 없다.

![image](https://github.com/yumin00/blog/assets/130362583/b01f4cd7-d7bf-4451-800a-fee61448e97a)
## Commit Log

event가 commit되면, 추가만 가능하고 변경은 불가능하다. 이때, 이벤트는 항상 로그 끝에 추가된다.

commit log에서 이벤트의 위치를 offset이라고 한다. 

## Offset
- offset은 증가만 하고 0으로 돌아가지 않는다
- 다른 consumer group에 속한 consumer들은 서로 관련이 없고, topic에 있는 event를 동시에 다른 위치해서 consume할 수 있다. 즉 서로 offset이 다를 수 있다.
  - 위 이미지처럼, Consumer Group A와 Consumer Group B의 consumer들은 서로 관련 없이 동시에 다른 offset에서 consume할 수 있다. 
- produce가 write하는 offset: LOG-END-OFFSET
- consumer가 read하고 처리한 후 commit한 offset: CURRENT-OFFSET
- LOG-END-OFFSET 와 CURRENT-OFFSET 의 차이 : Consumer Lag

# 2. Topic, Partition, Segment
topic, partition, segment 의 logical view 는 다음과 같이 표현할 수 있다.

![image](https://github.com/yumin00/blog/assets/130362583/72621784-1de4-4da9-a1c1-b28c69c45d34)

## Topic
메시지가 저장되는 장소, 토픽은 파티션으로 구성됨

## Partition
![image](https://github.com/yumin00/blog/assets/130362583/52a1720a-9143-4126-8abf-ec76a8577eb6)


- commit log
- 하나의 토픽은 하나의 파티션으로 구성
- 0부터 시작
- 토픽 내 파티션들은 독립적으로 작동
- 각 파티션의 offset은 서로 관련이 없음
- 병렬처리 향상(throughput 향상)을 위해 multi partition 사용을 권장
- 각 파티션은 브로커에 분산
- 파티션은 segment 파일로 구성됨

## Segment
![image](https://github.com/yumin00/blog/assets/130362583/212aa092-d037-4c3d-889c-dab0c9632fa3)

- 메시지가 저장되는 실제 물리 file
- 지정된 크키가 있어서, 실제 file보다 크거나 지정된 기간보다 오래되면 새 파일이 열리고 메시지는 새 파일에 추가됨
- 위 이미지에서 실제 메시지가 저장되는 곳은 segment 3이다.
- rolling 정책에 따라서 다음 파일이 열리는 기준을 정할 수 있다.
  - log.segment.bytes: default 1기가
  - log.roll.hours: default 168시간

## Broker
producer가 메시지를 프로듀스하면 받아서 저장하고, consumer가 메시지를 컨슘하면 메시지를 전달해주는 역할이다. 

브로커 노드 안에 파티션들이 만들어져 있는데, 파티션 당 오직 하나의 segment가 활성화되어 있다.

