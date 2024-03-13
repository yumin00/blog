---
title: "Kafka Topic에 대해 알아보자"
date: 2024-03-12T22:39:24+09:00
draft: true
categories :
- DataEngineering
- Kafka
---

# Kafka 주요 요소
Topic에 대해 알아보기 전, 카프카의 주요 요소에 대해 먼저 간단히 알아보자.

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
프로듀서와 컨슈머는 서로 완전히 분리되어 있기 때문에 누가 produce 하는지 누가 consume 하는지 서로 알 수 없다.

다른 consume group에 속한 consumer들은 서로 관련이 없고, topic에 있는 event를 동시에 다른 위치해서 consume할 수 있다. 즉 서로 offset이 다를 수 있다.

## Commit Log
event가 commit되면, 추가만 가능하고 변경은 불가능하다. 이때, 이벤트는 항상 로그 끝에 추가된다.

commit log에서 이벤트의 위치를 offset이라고 한다. 

## Offset
- offset은 증가만 하고 0으로 돌아가지 않는다
- produce가 write하는 offset: LOG-END-OFFSET
- consumer가 read하고 처리한 후 commit한 offset: CURRENT-OFFSET
- LOG-END-OFFSET 와 CURRENT-OFFSET 의 차이 : Consumer Lag

## Topic, Partition, Segment
### Topic
메시지가 저장되는 장소, 토픽은 파티션으로 구성됨

### Partition
- commit log
- 하나의 토픽은 하나의 파티션으로 구성
- 0부터 시작
- 토픽 내 파티션들은 독립적으로 작동
- 각 파티션의 offset은 서로 관련이 없음
- 병렬처리 향상(throughput 향상)을 위해 multi partition 사용을 권장

### Segement
- 메시지가 저장되는 실제 물리 file
- 지정된 크키가 있어서, 실제 file보다 크거나 지정된 기간보다 오래되면 새 파일이 열리고 메시지는 새 파일에 추가됨

## Broker
프로듀서가 메시지를 프로듀스하면 받아서 저장하고, 컨슈머가 메시지를 컨슘하면 메시지를 전달해주는 역할이다. 

브로커 노드 안에 파티션들이 만들어져 있는데, 파티션 당 오직 하나의 segment가 활성화되어 있다.

rolling 정책: log.segment.bytes(1기가) / log.roll.hours(168시간)