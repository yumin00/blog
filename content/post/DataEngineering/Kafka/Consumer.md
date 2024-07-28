---
title: "Kafka Consumer에 대해 알아보자"
date: 2024-03-27T20:19:22+09:00
draft: false
categories :
- DataEngineering
- Kafka
---
패스트캠퍼의 `한번에 끝내는 Kafka Ecosystem` 강의 내용을 바탕으로 Consumer를 정리해보고자 한다.

# Kafka Consumer
## Consumer
consumer는 kafka 에서 메시지를 읽는 주체이다. consumer들은 고유의 속도로 commit log로부터 순서대로 메시지를 읽는다.(poll)

서로 다른 consumer group의 consumer들은 서로 관련이 없기 때문에, 같은 파티션에서 메시지를 읽고 있더라고 아무런 관련이 없고, 동시에 다른 위치에서 읽을 수 있다.

## Consumer Offset
![image](https://github.com/yumin00/blog/assets/130362583/34761198-b99c-49e7-adf3-cf58503aab71)

consumer들은 메시지를 읽은 다음 읽은 위치를 표시하는데, 이를 offset이라고 한다. 데이터를 읽은 위치를 자동/수동으로 commit 함으로써 다시 읽는 것을 방지한다.

consumer가 메시지를 읽고 commit한 다음, 해당 offset을 `__consumer_offsets` 라는 internal topic에 저장한다. (group명 : 토픽이름 : 파티션 번호 : 그 다음 읽을 위치)

`__consumer_offsets` 에서는 consumer의 offset을 저장하여 관리하는 역할을 한다.


## Multi-Partitions With Single Consumer
여러 개의 파티션이 있지만 하나의 consumer만 존재한다면 어떻게 될까?

한 consumer가 모든 파티션에서 데이터를 consume할 뿐만 아니라, 각 파티션에서의 consumer offset을 별도로 기록해야 한다.

consumer가 하나임에도 불구하고 각각 다 기록하는 이유는 무엇일까?
운영을 하다보면, consumer가 하나로도 충분하다가도, consumer를 추가해야할 때가 발생한다. 그러면 파티션별로 오프셋을 구분해야하기 때문에 각 파티션마다 consumer offset을 기록함으로써 다른 consumer가 추가되더라도 문제 없이 운영이 가능하다.


## Consumer Group
![image](https://github.com/yumin00/blog/assets/130362583/e4cdfdbb-97b5-4f0a-a0c5-7e86e1029edc)
동일한 `group.id`로 구성된 consumer들은 하나의 consumer group을 형성한다. 파티션은 한 consumer group 내의 하나의 consumer에 의해서만 사용되어야 한다. 즉, 한 group 내의 consumer들이 같이 파티션을 사용할 수 없다.

4개의 파티션에 4개의 consumer라면 1:1로 분배하게 된다.

## message ordering(순서)
### 파티션이 여러 개일 경우
모든 메시지에 대해서 순서 보장이 불가능 하다.

### 파티션이 1개일 경우
모든 메시지에 대하여 순서 보장이 가능하다. 하지만 처리량이 저하된다는 단점이 있다.

파티션을 1개로 구성해서 모든 메시지에 대해서 전체 순서 보장을 해야하는 경우는 많지 않다. 대부분의 경우에는 key로 구분할 수 있는 메시지들의 순서 보장이 필요한 경우가 많기 때문에, 순서 보장이 필요할 때는 key를 잘 사용하면 좋다!

만약 운영 중에, 파티션 개수를 변경하면 순서를 보장할 수 업식 때문에, 이를 잘 생각하고 파티션 개수를 조절해야 한다.


## Key Cardinality
![image](https://github.com/yumin00/blog/assets/130362583/6ee76689-72f2-4d61-9f2a-97f0b00067ac)

cardinality란 특정 데이터 집합에서 유니크한 값의 개수를 말한다. 즉, 카프카에서는 한 토픽의 여러 개의 파티션에 있는 데이터의 개수라고 할 수 있다. 위의 사진처럼 데이터의 개수가 적은 파티션의 consumer들은 놀게 된다.

즉, 파티션의 메시지들은 Key를 통해 분배되기 때문에 Key Cardinality는 consumer group의 개별 consumer가 수행하는 작업의 양에 영향을 준다. 

카프카는 메세지 분포를 하여 병렬처리를 하는 것이 목적인데, 키로 인해 특정 consumer만 일해서 느리게 처리되면 좋지 않은 상황이다. 즉, 분포를 제대로 하지 않으면 consumer의 워크 로드가 고르게 분포되지 않아 사이드 이팩트가 발생할 수 있다.

key는 json, avro 등 여러 필드가 있는 복잡한 객체로 만들어 되고, key 를 잘 분포하여 consumer 모두가 일할 수 있도록 하는 것이 중요하다.

## Consuemr Rebalancing
만약 4개의 파티션이 있고 4개의 consumer가 있는 consumer group이 있다고 가정해보자. 그러면 파티션과 consumer는 1:1 매칭이 되어 각자 메세지를 컨슘할 것이다.

만약, 한 consumer에 오류가 발생하여 사라졌다면, 남아있는 consumer 중 하나가 해당 파티션의 메세지도 같이 컨슘하게 된다. 이를 Consuemr Rebalancing이라고 한다.

### consumer load balancing
로드밸런싱 된다. 한 파티션 당 하나의 컨슈머

### Partition assignment
파티션을 컨슈머 쪽으로 어떻게 할당할거냐? 

### consumer group coordination
그룹 코디네이터는 하나의 브로커가 담당함 코드네이터와 그룹 리더가 상호작용하여 어떤 파티션에 매핑할지 결정함.

consumer_offsets이라는 토픽에는 디폴트로    50개의 파티션이 있는데, 이 파티션들도 브로커에 분산되어 있다.

컨슈머 그룹내의 컨슈머들의 모든 오프셋은 컨슈머 오프셋 토픽의 하나의 파티션에 저장이되는데, 해당 파티션의 리더를 가지고 잇는 브로커가 코더네이터가 됨.

왜 그룹 코디네이터가 안하고, 그룹 리더가 결정하지?

컨슈머 리밸런싱: 파티션 할당받고,,등등 다시 일어나는 거를 리밸런싱

### 컨슈머 하트비츠
max.poll.interval.ms
하트비트는 보내는데, 폴이 너무 오래 걸리면, 5분을 넘어서면 문제가 있다고 판단함.

### 과도한 리밸런싱을 피하는 방법 (성능 최적화)
rejoin: 특정 그룹 인스턴스 아이디를 가지고 있는 클라이어트가 빠졌다가 다시 붙을 때는 리밸런스 안하고 그냥 붙게 해줌

session.timeout.ms를 gorup.min.session.timeout.ms와 group.max.session.timout 사이값으로 설정하면 좋음

너무 크면 장애 인지를 못함 (max.poo..interval.ms)

성능 테스트를 통해 적정값을 찾아가는 것이 중요함.



# 파티션 어사인먼트 스트레티지
같은 딜리버리 키를 가지고 데이터를 땡겨갈 수 있음.

스티키
: 기존 할당을 유지하고 나머지 부분만 재할당


# consume rebalancing process(파티션 할당 과정)
이때 리더는 제일 먼저 joinGroup 요청ㅇ을 보낸 컨슈머가 리더

## eager rebalancing 프로토콜
조인그룹을 요청하는 순간 파티션 취소하고 컨슘을 멈춤
할당이 이뤄져야 다시 할당

## incremental cooperative rebalancing protocol
정말 빠질 파티션만 revoke

실제 구현 ㅏ방식: copperative sticky assignor