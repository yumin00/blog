---
title: "Redis에 대해 알아보자"
date: 2024-06-24T21:13:34+09:00
draft: false
categories :
- DataEngineering
---

## Redis 배경
Redis는 "Remote Dictionary Server"의 약자로, 인메모리 데이터베이스이다. Redis는 성능, 확장성, 유연성에 대한 요구에서부터 시작되었다.

기존 RDBMS는 디스크 기반의 데이터베이스이기 때문에 높은 I/O 비용으로 성능 문제가 발생할 수 있었다. Redis는 메모리 내에서 모든 데이터를 처리하기 때문에 RDBMS보다 더 빠른 읽기/쓰기 성능을 제공한다.

## Redis 특징
### 복잡한 데이터 구조 지원
Redis는 key-value 구조를 갖는 데이터베이스이다. 단순히 key-value를 넘어서 다양한 데이터 구조(문자열, 리스트, 해시, 셋, 정렬된 셋, 비트맵)를 지원하여 다양한 데이터 구조를 활용할 수 있도록 해준다.

### 데이터 백업
Redis는 인메모리 데이터베이스이기 때문에 속도가 빠르다는 장점이 있지만, 메모리 특성상 저장된 데이터가 휘발될 가능성이 있다. Redis는 이를 해결하고자 관리하고 있는 데이터에 영속성을 제공하기 위해 메모리에 있는 데이터를 디스크에 백업하는 기능을 제공한다.

#### RDB 방식
Redis DataBase 방식은 메모리에 있는 데이터 전체에서 스냅샷을 작성하고, 이를 디스크에 저장하는 방식이다. 하지만, 스냅샷 이후에 변경된 데이터는 복구할 수 없다는 단점이 있다.

#### AOF 방식
Append Only File 방식은 데이터가 변경되는 이벤트가 발생하면 이를 모두 로그에 저장하는 방식이다. 모든 데이터의 변경 사항을 기록하기 때문에 RDB 방식보다 데이터 유실량이 적지만 모든 로그를 가지고 있기 때문에 RDB 방식보다 로딩 속도가 느리고 파일 크기가 큰 것이 단점이다.

## Redis 싱글 스레드
Redis는 기본적으로 싱글 스레드로 작동한다. 

### 싱글 스레드 모델
- 이벤트 루프 기반: Redis는 싱글 스레드 이벤트 루프를 사용하여 클라이언트의 요청을 처리한다. 즉, 사용자의 요청이 들어오면 이는 Event Queue에 적재하고 하나씩 처리한다. 모든 요청은 순차적으로 처리된다.
- 락이 필요 없음: 싱글 스레드 모델이기 때문에 여러 스레드가 동시에 접근하지 않아 데이터 일관성을 위해 락을 걸 필요가 없다.
- 빠른 처리 속도: 메모리 내에서 작업이 처리되기 때문에 싱글 스레드로 동작하더라도 작업을 효과적으로 처리할 수 있다.

### 성능 최적화
Redis는 싱글 스레드를 통해 높은 성능을 제공하기 위해서 여러가지 최적화 기법을 사용하고 있다.

- 입출력 다중화: Redis는 epoll, kqueue, select 와 같이 I/O 다중화 기술을 사용하여 많은 클라이언트의 연결을 효율적으로 처리한다.
- 파이프라이닝: 클라이언트는 여러 명령어를 파이프라인으로 묶어 한 번에 보낼 수 있고, 이는 네트워크 왕복 시간을 줄여 성능을 향상시킨다.

### 멀티 코어
Redis는 기본적으로 싱글 스레드로 작동하지만, 성능을 극대화하기 위해서 멀티 코어로 사용할 수 있다.

- 멀티플 Redis 인스턴스: 하나의 서버에서 여러 Redis 인스턴스를 실행하여 각 인스턴스가 다른 CPU 코어에서 작동하게 하는 방법이다. 수평적 확장을 통해 전체 시스템의 처리량을 증가시킬 수 있다.
- Redis 클러스터: 데이터를 여러 노드에 분산 저장하는 Redis 클러스터를 사용하여 수평적으로 확장할 수 있다. 각 노드는 독립적으로 작동하고, 전체 클러스터의 처리량을 크게 향상시킬 수 있다.

## Redis 동작방식
![Image](https://github.com/user-attachments/assets/ed966908-9124-49d7-a7f1-7c2b2ce5d07e)
1. 클라이언트가 Redis 서버에 연결 요청
- Redis는 클라이언트의 TCP 연결을 수락하고, 해당 소켓을 epoll에 등록
2. 클라이언트가 Redis 서버에 명령어 전송
- SET key value, GET key 등의 명령을 소켓을 통해 전달
3. Redis의 epoll이 요청을 감지
- 이벤트 루프(Event Loop)가 실행되면서 epoll_wait() 호출
- epoll_wait()은 모든 등록된 소켓을 감시하며 이벤트가 발생하면 반환
4. 이벤트 큐에 요청 추가(필요한 경우)
- 블로킹 가능성이 있는 요청은 이벤트 큐에 추가
- 단순한 I/O 작업은 바로 실행
5. 단일 스레드가 요청 처리
- 이벤트 루프가 요청을 하나씩 가져와 명령 실행
- GET key → 데이터 조회, SET key value → 데이터 저장 후 응답 버퍼에 저장
6. 응답 버퍼에 직렬화된 데이터 저장
- Redis 프로토콜(RESP) 형식으로 직렬화
7. I/O 스레드가 응답을 클라이언트에게 전송
- 여러 클라이언트가 있더라도 병렬로 응답 가능

## Redis Architecture
### Replication 아키텍처
![image](https://github.com/yumin00/blog/assets/130362583/f15b5301-46b0-4793-ab63-09cbcb790f79)

Replication 아키텍처는 Redis 데이터베이스의 데이터를 여러 복제본으로 유지하여 고가용성과 읽기 성능을 향상시키는 데 사용된다.

#### 특징
- 마스터-슬레이브 구조: 하나의 마스터 노드와 하나 이상의 슬레이브 노드로 구성된다. 마스터 노드는 쓰기 작업을 처리하고, 슬레이브 노드는 마스터의 데이터를 복제하여 읽기 작업을 분산 처리한다.
- 비동기 복제: 슬레이브 노드는 비동기적으로 마스터 노드의 데이터를 복제한다. 이는 데이터 일관성을 유지하면서도 복제 속도를 높인다.
- 읽기 성능 향상: 슬레이브 노드는 읽기 요청을 처리할 수 있으므로, 여러 슬레이브 노드를 두어 읽기 성능을 크게 향상시킬 수 있다.
- 장애 복구: 마스터 노드에 장애가 발생하면 슬레이브 노드를 승격하여 새로운 마스터 노드로 사용할 수 있다.
- 고가용성 (High Availability, HA) 제공

해당 특징을 보면, 여러 슬레이브 노드를 두어 읽기 성능을 크게 향상시킬 수 있다고 한다. 그런데 Redis는 기본적으로 싱글 스레드 모델이기 때문에 한 번에 하나의 작업만 처리할 수 있다.
그러면 여러 개의 슬레이브 노드를 둔다는 것은 멀티 코어를 사용한다는 것인지, 혹은 각각의 노드가 독립적인 스레드라는 것일까?

![image](https://github.com/yumin00/blog/assets/130362583/a8479325-c63c-4bc2-9333-480591d160c5)
멀티 코어는 하나의 프로세스에서 여러 개의 코어를 사용하는 것이다. 하지만 Replication 아키텍처는 하나의 노드가 하나의 Redis 인스턴스로, 애초에 여러 개의 프로세스가 있고 각 노드가 각각의 코어를 갖고 있는 것이다!

#### 동작 방식
1. 마스터 노드에 쓰기 작업이 발생하면, 해당 변경 사항이 슬레이브 노드로 전파된다.
2. 슬레이브 노드는 마스터 노드의 변경 사항을 비동기적으로 수신하고, 자신의 데이터베이스를 업데이트한다.
3. 클라이언트는 마스터 노드에 쓰기 요청을 보내고, 슬레이브 노드에 읽기 요청을 보낸다.

### Sentinel 아키텍처
![image](https://github.com/yumin00/blog/assets/130362583/c515dde2-a8b5-4bd4-888a-3b9467cbff57)
Sentinel 아키텍처는 자동 장애 조치(failover) 기능을 제공하는 Redis 감시 시스템이다. Sentinel은 Redis 서버를 모니터링하고 장애가 발생하면 자동으로 장애를 조치힌다.

#### 특징
- 모니터링: Sentinel은 마스터 노드와 슬레이브 노드를 지속적으로 모니터링하며 노드의 가용성을 확인한다.
- 자동 장애 조치: 마스터 노드에 장애가 발생하면, 슬레이브 노드 중 하나를 마스터 노드로 승격시킨다.
- 알림: 노드 상태 변화, 장애 조치 등 주요 이벤트에 대해 관리자에게 알림을 보낸다.
- 구성 관리: 새로운 마스터나 슬레이브 노드의 구성을 자동으로 업데이트하여 클러스터의 일관성을 유지한다.
- Sentinel 노드도 장애 상황이 발생할 수 있기 때문에 반드시 3대 이상의 홀수로 존재해야 한다.
- 많은 리소스가 필요하므로 Sentinel과 마스터 노드 혹은 슬레이브 노드를 같은 서버에 올려 사용하기도 한다.

#### 동작방식
1. 여러 Sentinel 인스턴스가 Redis 서버를 모니터링한다.
2. 노드가 응답하지 않으면 장애를 감지한다.
3. 장애가 감지되면, 투표를 통해 새로운 마스터 노드를 선출하고 슬레이브 노드를 새로운 마스터로 승격시킨다.
4. 클러스터 구성 정보가 자동으로 업데이트 되고, 클라이언트는 새로운 마스터 노드로 전환된다.

### Cluster 아키텍처
Cluster 아키텍처는 데이터 분산과 수평 확장을 위해 설계된 Redis의 분산 데이터베이스 솔루션이다. 대규모 데이터 셋을 처리하고, 고갸용성을 제공하는데 적합하다. Redis 3.0 버전 이후부터 제공된다.

#### 특징
![image](https://github.com/yumin00/blog/assets/130362583/358d95c7-c16b-4895-81ae-81dcabb7d7e4)
- 데이터 분산: 데이터를 여러 노드에 분산 저장하고, 각 노드가 전체 데이터의 일부만 저장한다. 이때 "슬롯" 단위로 나누어 각 노드에 할당하는 방식으로 구현된다.
> 슬롯
> 
> Redis는 해시 슬롯을 사용하여 데이터를 분산 저장한다. 각 키는 CRC16 해시 함수에 의해 0부터 16383 사이의 슬롯 번호로 변환된다. 이 슬롯 번호는 클러스터의 특정 노드에 할당된다.
> 
> 예시)
> 
> `user:1000`이라는 키가 있다고 가정해보자. 이 키는 CRC16 해시 함수에 의해 해싱되어 16384개의 슬롯 중 하나에 매핑되는데 예를 들어 슬롯 5471에 매핑된다고 가정해보자.
> 클러스터 내의 각 노드는 특정 범위의 슬롯을 담당한다. 예를 들어, 노드 A가 슬롯 0-5500, 노드 B가 슬롯 5501-11000, 노드 C가 슬롯 11001-16383을 담당한다고 가정하면 `user:1000` 는 노드 A에 저장되는 것이다.

- 고가용성: 각 노드는 다른 노드와 데이터를 복제하여 고가용성을 보장한다. 하나의 노드에 장애가 발생해도, 다른 노드가 해당 데이터를 유지하고 있다.
- 자동 장애 조치: 클러스터 내에서 장애가 발생하면, 남은 노드들이 자동으로 장애 조치를 수행하고 클러스터를 복구한다.
- 수평 확장: 노드를 추가하거나 제거하여 클러스터의 용량과 성능을 쉽게 확장할 수 있다.

#### 동작 방식
1. 클러스터 내의 모든 데이터는 16384개의 슬롯으로 나누어지고, 각 슬롯은 특정 노드에 할당된다.
2. 클러스터는 요청을 보낼 때, 키 해싱을 통해 해당 키를 저장할 슬롯을 결정하고 해당 슬롯을 담당하는 노드에 요청을 보낸다.
3. 노드 간에 데이터를 복제하여, 노드 장애 시 자동으로 복구할 수 있다.
4. 클러스터 상태는 지속적으로 노드 간에 교환되어, 클러스터의 일관성을 유지한다.

## Redis 사용 용도
### 캐싱
Redis는 주로 캐싱에 사용된다. 자주 조회되는 데이터를 메모리에 저장하여 성능을 높이는 데 유용하다.

### 세션 관리
웹 애플리케이션에서 사용자 세션 정보를 저장하고 관리하는 데 Redis를 사용할 수 있다. 세션 정보를 빠른 속도로 실시간으로 공유해야 하는 환경에서 유용하다.

### 실시간 데이터 분석
Redis는 속도가 빠르기 때문에 실시간 데이터 처리와 분석에서 유용하다. 

### 메세지 브로커
Pub/Sub 모델을 지원하므로 메세지 브로커로 사용할 수 있다. 즉, 서비스 간 비동기 메시징 처리에 사용될 수 있다.

## Redis 사용사례
실제 기업에서 Redis를 어떤식으로 사용하고 있는지에 대해 스터디해보았다.

### 우아한 형제들 ([참고](https://techblog.woowahan.com/2709/))
배달의민족에서는 사용자에게 선물하기 기능을 제공하고 있다. 이때 상품에 대한 재고관리 요구사항이 존재한다고 한다.

```
1. 상품의 권종별로 전체 재고수량과 인당 재고수량이 관리되어야 한다.
2. 상품은 전체재고량을 초과하여 판매되면 안된다.
3. 판매가 시작된 상품의 전체 재고수량은 감소시킬 수 없다.
```

우아한형제들에서는 해당 요구사항을 충족시키기 위해서, 전체 재고량 관리와 재고 사용량을 분리하여 저장하였다. 재고 사용량 증가와 감소는 동시성 이슈를 없애기 위해 단일 스레드인 Redis를 사용하면서, 재고 사용량의 데이터가 유실되면 안 되기 때문에 RDB에 싱크할 수 있도록 했다고 한다.

이러한 설계를 통해 실제 상품권 구매에 대한 트랜잭션은 단일 스레드는 Redis를 통해 동시성 이슈를 처리하였고, RDB를 사용하여 데이터 유실을 방지한 것을 확인할 수 있다!

해당 케이스를 통해, 확실한 트랜잭션 관리가 필요한 경우에는 Redis를 사용함으로써 트랜잭션도 관리하고 속도도 빠르게 진행할 수 있다는 것을 알 수 있었다.