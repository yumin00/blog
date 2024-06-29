---
title: "DB에서 동시성 관리를 어떻게 할까?"
date: 2024-05-26T23:21:55+09:00
draft: false
categories :
- DataBase
- DataEngineering
---

데이터베이스에는 한 번에 많은 요청이 들어올 수 있다. 이러한 경우에 동시성 관리를 하지 않는다면, 문제가 발생할 수 있다. 데이터의 무결성을 유지하기 위해서 동시성 관리는 아주 중요하다. 데이터베이스에서는 동시성을 어떻게 관리하는지에 대해 공부해 보고자 한다.

## 잠금 (Locking)
잠금이란, 트랜잭션이 잠금을 설정했다면 해제할 때까지 데이터를 독점적으로 사용하여 다른 트랜잭션은 해당 데이터에 접근할 수 없는 것을 의미한다.

### 종류
- 공유 잠금 (Shared Locking): 트랜잭션이 특정 데이터에 대해 공유 잠금을 설정하면 해당 트랜잭션은 읽기만 가능하고, 다른 트랜잭션도 데이터 읽기가 가능하다.
- 전용 잠금 (Exclusive Locking): 트랜잭션이 특정 데이터에 대해 전용 잠금을 설정하면 해당 트랜잭션을 읽기/쓰기가 가능하고, 다른 트랜잭션은 읽기/쓰기가 모두 불가능하다.

### 블로킹 현상 (Blocking)
만약 A 트랜잭션이 name 테이블의 첫 번째 row를 수정하기 위해 전용 잠금을 설정했다고 가정해보자. name 테이블의 첫 번째 row를 읽어야하는 B 트랜잭션은 A 트랜잭션이 모두 완료할 때까지 기다려야 하는데, 이렇게 기다리는 현상을 블로킹 현상이라고 한다.

블로킹 현상은 어플리케이션의 성능에 좋지 않은 영향을 미치기 때문에 블로킹을 최소화는 것이 중요하다.

#### 블로킹 현상 완화 전략
블로킹 현상을 완화하기 위해서는 다음과 같은 전략을 세울 수 있다.

- 트랜잭션 내에서 잠금을 유지하는 시간을 최소화
- 트랜잭셔을 작게 분할하여 블로킹 현상을 완화

### 교착 상태 (DeadLock)
만약 A 트랜잭션이 name 테이블의 첫 번째 row를 수정한 뒤, age 테이블의 첫 번째 row를 수정한다고 가정해보자. 그리고 B 트랜잭션을 age 테이블의 첫 번째 row를 수정하고, name 테이블의 첫 번째 row를 수정한다.

이때, 각 트랜잭션이 전용 잠금을 사용하여 첫 번째 수정을 마친 다음, 두 번째 작업을 진행하려고 할 때 서로 수정해야하 하는 테이블이 Lock 상태로 서로의 트랜잭션이 끝날 때까지 대기하기 때문에 교착 상태에 빠질 수 있다.

## 낙관적 동시성 제어 (Optimistic Concurrency Control)
낙관적 동시성 제어란, 충돌이 많이 발생하지 않을 것이라고 가정하고 동시성을 제어하는 방식이다. 주로 읽기 작업이 많고, 쓰기 작업이 적은 환경에서 효과적이다. 낙관적 동시성 제어는 다음과 같은 단계로 동작한다.

### 동작 방식
#### 1. 트랜잭션 시작 및 작업 수행
트랜잭션은 필요한 읽기 및 계산 작업을 수행한다. 이 단계에서 실제로 데이터베이스의 데이터는 수정되지 않고, 변경될 데이터를 로컬 변수 혹은 임시 데이터 구조에 저장한다.

#### 2. 결과 검증
트랜잭션이 종료되면, 데이터베이스는 해당 트랜잭션이 실행되는 동안 다른 트랜잭션에 의해 데이터베이스의 데이터가 변경되었는지 검증한다. 이때 검증 방식은 여러가지가 있다.

- 타임스탬프 비교: 각 데이터에 타임스탬프를 부여하여 트랜잭션이 시작된 이후에 변경된 데이터가 있는지 확인
- 버전 관리: 각 데이터에 버전을 부여하여 트랜잭션이 시작된 이후에 버전이 변경되었는지 확인

#### 3. 커밋 or 롤백
검증 단계에서 충돌이 발생하지 않으면 데이터베이스에 변경 사항을 커밋한다.

만약 충돌이 발생했다면 트랜잭션은 롤백되고 재시도된다.

## 비관적 동시성 제어 (Pessimistic Concurrency Control)
비관적 동시성 제어란, 충돌이 많이 발생할 것이라고 가정하고 동시성을 제어하는 방식이다. 데이터에 접근하기 전에 미리 잠금을 걸고, 다른 트랜잭션이 접근할 수 없도록 하는 방식이다. 이 방식은 주로 쓰기 작업이 많은 환경에서 효과적이다.

### 동작 방식
#### 1. 잠금 획득
트랜잭션이 특정 데이터에 접근하고자 하면 먼저 해당 데이터에 대한 잠금을 획득해야 한다.

#### 2. 데이터 접근
잠금을 획득하면 데이터에 접근하여 작업을 수행한다. 잠금이 유지되는 동안에는 다른 트랜잭션은 해당 데이터에 접근할 수 없다.

#### 3. 잠금 해제
트랜잭션이 완료되면 잠금을 해제한다. 잠금 해제는 트랜잭션이 커밋되었거나 롤백되었을 때 이루어진다.

#### 4. 교착 상태 처리
비관적 동시성 제어는 서로의 잠금이 풀리기를 기다리는 교착 상태를 처리해야 한다. 교착 상태 처리 방법은 여러가지가 있다.

- 타임아웃: 일정 시간 동안 잠금을 획득하지 못하면 트랜잭션을 롤백한다.
- 교착 상태 감지: 주기적으로 교착 상태를 감지하고, 교착 상태에 있는 트랜잭션 중 하나를 롤백한다.
- 교착 상태 예방: 트랜잭션이 잠금을 획득하는 순서를 정하여 교착 상태가 발생하지 않도록 예방한다.

## 다중 버전 동시성 제어(Multi-Version Concurrency Control, MVCC)
MVCC란, 각 데이터 항목의 여러 버전을 유지함으로써 트랜잭션 간의 충돌을 피하는 방식이다. 주로 읽기 작업이 많은 환경에서 높은 성능을 제공한다.

### 데이터 버전 관리
MVCC는 데이터의 여러 버전을 유지한다. 각 버전은 다음과 같은 메타데이터를 포함한다.

- 트랜잭션 타임스탬프 or ID: 해당 버전을 수정하거나 생성한 트랜잭션의 타임스탬프 또는 ID
- 유효기간: 각 데이터 버전의 시작 타임스탬프와 종료 타임스탬프를 나타내며, 해당 버전이 유효한 기간을 나타낸다.

### 트랜잭션 읽기
트랜잭션이 데이터 항목을 읽을 때, MVCC는 트랜잭션 타임스탬프에 해당하는 데이터 버전을 제공한다. 트랜잭션이 시작된 시점의 스냅샷을 기반으로 데이터를 읽기 때문에, 다른 트랜잭션이 데이터 항목을 수정하더라도 영향을 받지 않는다.

### 트랜잭션 쓰기
트랜잭션이 데이터를 수정할 때, MVCC는 기존 데이터를 덮어쓰지 않고 새로운 버전을 생성한다. 새로운 버전은 트랜잭션 타임스탬프와 함께 저장된다. 이는 다른 트랜잭션이 이전 버전을 읽을 수 있게 하고, 데이터의 일관성을 유지한다.

### 장점
- 트랜잭션이 서로 독립적으로 데이터를 읽고 쓸 수 있어 높은 동시성을 제공한다.
- 데이터 읽기 작업에서 잠금을 사용하지 않기 때문에 읽기 성능이 우수하다.
- 각 트랜잭션은 자신만의 일관된 스냅샷을 볼 수 있기 때문에 데이터 일관성이 보장된다.

### 단점
- 여러 버전을 저장해야 하기 때문에 저장소 사용량이 증가한다.
- 데이터를 수정할 때마다 새로운 버전을 생성해야하기 때문에 쓰기 성능이 저하될 수 있다.

## 직렬성 (Serialization)
직렬성는 트랜잭션이 병렬로 실행되더라도 결과가 마치 트랜잭션이 순차적으로 실행된 것처럼 보이도록 하는 방법이다. 이는 가장 강력한 동시성 제어 방식으로, 트랜잭션 간의 모든 충돌을 방지한다. 이를 통해 데이터 무결성을 유지할 수 있다.

### 시리얼 가능성 (Serializability)
시리얼 가능성은 트랜잭션 스케줄링의 기준으로, 트랜잭션들이 동시에 실행될 때의 결과가 특정 순서로 순차적으로 실행된 것과 동일한 경우를 말한다.
시리얼 가능한 스케줄은 데이터의 일관성을 보장합니다.

### 구현방법
#### 2단계 잠금 프로토콜(Two-Phase Locking, 2PL)
2PL은 시리얼 가능성을 보장하는 가장 널리 사용되는 방법이다. 이 프로토콜은 두 가지 단계로 이루어져 있다.

- 확장 단계(Growing Phase): 트랜잭션이 필요한 모든 잠금을 획득하지만, 잠금을 해제하지 않는다.
- 축소 단계(Shrinking Phase): 트랜잭션이 더 이상 잠금을 획득하지 않고, 기존 잠금을 해제한다.

2PL은 트랜잭션이 확장 단계에서 잠금을 획득하고, 축소 단계에서 잠금을 해제함으로써 시리얼 가능성을 보장한다.

#### 타임스탬프 순서
타임스탬프 순서 기법은 각 트랜잭션에 고유한 타임스탬프를 부여하고, 트랜잭션의 순서를 타임스탬프에 따라 결정하는 방법이다.

- 각 데이터는 읽기 타임스탬프와 쓰기 타임스탬프를 가진다.
- 트랜잭션은 자신의 타임스탬프보다 최신의 타임스탬프를 가진 데이터는 접근할 수 없다.
- 트랜잭션이 충돌이 일으키면, 타임스탬프를 기준으로 충돌을 해결하고 트랜잭션을 재시도하거나 중단한다.

#### 직렬화 가능한 스냅샷 격리(Serializable Snapshot Isolation, SSI)
> 스냅샷 격리 (Snaphsot Isolation)
> 
> 스냅샷 격리(Snapshot Isolation, SI)는 데이터베이스 관리 시스템에서 사용되는 트랜잭션 격리 수준 중 하나로, 트랜잭션이 자신만의 데이터베이스 스냅샷을 사용하여 일관된 읽기를 수행할 수 있도록 한다.
> 이는 주로 다중 버전 동시성 제어(Multi-Version Concurrency Control, MVCC)를 기반으로 구현된다.

SSI는 스냅샷 격리를 확장하여 시리얼 가능성을 보장하는 방법이다. 각 트랜잭션은 자신만의 데이터 스냅샷을 사용하여 동작하며, 데이터 충돌을 감지하고 해결한다. SSI는 다음과 같은 방식으로 동작한다.

- 스냅샷 생성: 트랜잭션이 시작될 때 데이터베이스의 일관된 스냅샷을 생성한다.
- 충돌 감지: 트랜잭션이 데이터를 읽거나 쓸 때 충돌을 감지한다. 충돌이 감지되면, 트랜잭션을 중단하거나 재시도한다.
- 충돌 해결: 충돌이 발생한 트랜잭션 중 하나를 선택하여 중단하고, 나머지 트랜잭션을 계속 진행한다.

### 장점
- 트랜잭션이 직렬적으로 실행된 것과 같은 결과를 보장하여 데이터 일관성을 유지한다.
- 동시에 실행되는 트랜잭션 간의 충돌을 효과적으로 방지한다.

### 단점
- 많은 트랜잭션이 동시에 실행될 때, 대기 시간이 증가하고 시스템 처리량이 감소할 수 있다.
- 트랜잭션이 잠금을 오래 유지하면 다른 트랜잭션이 대기해야 하므로, 전체 시스템의 효율성이 떨어질 수 있다.