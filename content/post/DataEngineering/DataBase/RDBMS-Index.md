---
title: "PostgreSQL의 Index에 대해 알아보자"
date: 2024-04-18T19:52:46+09:00
draft: false
categories :
- DataBase
- DataEngineering
---

RDMBS에서는 Index를 사용한다. 그렇다면 RDBMS에서는 왜 index를 사용할까?

# Index란?
인덱스란, 색인/지표를 의미하는 단어로 테이블의 검색 속도를 향상시키기 위해 사용하는 자료구조이다. RDBMS에서도 테이블 검색 속도를 향상시키기 위해 인덱스를 사용한다고 한다.

# 인덱스를 사용하면 왜 속도가 빨라질까?
예를 들어 백과사전에서 원하는 단어를 찾아야할 때 모든 단어를 다 훑어보면서 단어를 찾으려고 한다면 시간이 아주 오래 걸릴 것이다. 그래서 백과사전에는 색인이 존재하여 단어를 좀 더 쉽고 빠르게 찾을 수 있게 해준다.

이와 같이 테이블에서도 원하는 데이터를 찾을 때, 모든 데이터를 검색하여 찾는다면 시간이 아주 오래 걸릴 것이다. 때문에, 백과사전의 색인처럼 Index를 사용하는 것이다.

![image](https://github.com/yumin00/blog/assets/130362583/66a4ae98-cea4-4eaa-b8db-59357c701992)

예를 들어서, 위 사진처럼 user_id를 Index로 갖는다고 가정해보자. 그러면 index는 user_id의 값과 해당 데이터의 포인터를 가지고 있다. 특정 user_id를 통해 score를 검색한다면,
index를 통해 먼저 user_id의 pointer를 얻고 해당 pointer를 통해 score를 찾을 수 있는 것이다.

이러한 방법을 통해서, 모든 데이터를 검색하지 않고도 index를 통해 데이터를 빠르게 찾을 수 있다.

인덱스를 구현하는 방법은 여러가지가 있는데, 대표적인 방법에 대해 알아보자.

# Hash Table
![image](https://github.com/yumin00/blog/assets/130362583/0902e423-a4ea-4451-9800-ff2816c01997)
해시 테이블은 (key, value) 로 데이터를 저장하는 자료구조이다. 

위 그림처럼, key를 넣으면 hash function을 통해 value를 얻음으로써 index의 값을 넣으면 해당 데이터의 주소를 가지고 데이터를 검색하는 자료구조인 것이다.

해시 테이블의 시간 복잡도는 O(1)로 매우 빠른 검색을 지원한다.

하지만 hash table은 사실 Index에는 많이 사용되지 않는다. 그 이유는 hash function은 정확히 일치하는 값만 찾을 수 있기 때문에,
LIKE나 부등호를 사용하는 쿼리 검색에는 용이하지 않기 때문이다.

# B+Tree
![image](https://github.com/yumin00/blog/assets/130362583/3276b27c-6d7b-42d6-beb5-8d90ab8ab082)

B+Tree는 트리 기반 구조로, DB 성능을 최적화하기 위해 설계된 것이다.

- 내부 노드: key / 리프 노드의 pointer
- 리프 노드: data의 pointer

해당 키를 찾을 때까지 내부 노드를 돌고, 키를 찾으면 리프 노드의 포인터를 통해 데이터 검색하는 방식으로 동작한다.

실제 데이터 포인터가 리프 노드에만 저장되어 있기 때문에, 리프 노드에 도달할 때까지 실제 데이터에 접근할 필요가 없기 때문에 불필요한 디스크 I/O를 줄일 수 있다.

또한, 리프 노드가 모두 같은 깊이에 위치하기 때문에, 어떤 데이터에 접근하든 접근 시간이 일정하다.

# Full Table Scan
Full Table Scan은 DB에서 데이터를 검색하는 방법 중 하나이다.

이는, 인덱스를 사용하지 않고 테이블의 모든 행을 처음부터 끝까지 순차적으로 읽고 데이터를 검색하는 방식이다.

테이블의 데이터 양이 적거나, 거의 모든 데이터를 반환해야하는 경우에는 Full Table Scan을 사용하는 것이 효율적일 수 있다.

또한, B+Tree는 leaf node까지 내려가서 데이터를 찾아야하기 때문에, 데이터 양이 적다면 Full  Table Scan을 사용하는 것이 더 빠를 수 있다.

# 인덱스의 단점
- 인덱스를 관리하기 위해서는 DB의 약 10% 정도의 공간이 필요하다는 단점이 있다.
- 데이터 삽입, 삭제, 수정이 이루어질 경우, index도 함께 업데이트 되어야 한다. 이로 인해 데이터 변경 작업이 느려질 수 있다.
  - 자주 변경되는 데이터에 대해 인덱스가 설정되어 있다면, 데이터 변경과 인덱스 업데이트가 같이 이루어져 느려질 수 있다.

# PostrgreSQL 에서의 Index
## PRIMARY KEY와 INDEX
실제로, RDBMS 중 하나인 postgreSQL에서는 Index를 어떻게 사용하고 있을지에 대해 알아보자.
기존에 사용하고 있던 테이블의 index를 확인해보았다.

<img width="345" alt="image" src="https://github.com/yumin00/blog/assets/130362583/aa78064d-1f37-416d-9c26-b210bffa6592">

해당 테이블은 `order` 테이블로, `id` column을 primary key 로 설정하였다. primary key로 설정한 `id` column이 자동으로 index에 생성된 것을 확인할 수 있다.
기본키는 테이블 내에서 고유한 데이터로, 데이터 조회나 참조가 자주 이루어지기 때문에, postgresql에서는 기본키는 자동으로 인덱스를 생성해주는 것 같다. 데이터베이스트의 성능 향상을 위함이 아닐까 생각한다.

## INDEX 설정 방법
postgresql에서는 기본키를 설정해주면 자동으로 index를 생성할 수 있다. 해당 방법만 아니라, 조회를 자주 하는 컬럼을 기준으로 직접 index를 생성할 수 있다.

<img width="1427" alt="image" src="https://github.com/yumin00/blog/assets/130362583/83df350f-53ea-4ec1-956e-459785d579a8">

혹은 SQL문을 사용하여 인덱스를 생성할 수도 있다.

```sql
CREATE INDEX sales_id_index ON order (sales_id);
```

## [공식문서](https://www.postgresql.org/docs/current/indexes-intro.html)
공식 문서를 읽어보면, postgresql에서는 인덱스를 통한 조회 쿼리가 발생했을 때 인덱스를 사용한 조회가 빠를지 혹은 인덱스 사용 없이 조회하는 것이 더 빠를지 판단한 후 데이터 조회를 진행한다고 한다.
문서에서는 `ANALYZE` 명령어를 통해 해당 판단이 제대로 이루어지고 있는지에 체크해보는 것을 권장하고 있다.

데이터를 조회할 때 postgresql에서는 쿼리 플래너(Query Planner)와 옵티마이저(Optimizer)를 사용한다고 한다. 그리고 다음 내용을 바탕으로 인덱스를 사용 여부를 결정한다고 한다.

- 통계 정보: 각 테이블과 컬럼별 데이터 분포, 행의 개수, 평균 행 길이 등의 정보를 분석
- 쿼리 조건: 쿼리 내 조건을 고려하여, 특정 행만 조회하는 경우에는 인덱스를 선택하고, 조건이 없거나 대부분의 행을 조회해야 하는 경우 전체 스캔
- 인덱스가 지원하는 연산 유형을 고려하여 B-트리, GiST, GIN 등 최적화된 인덱스 타입 선택

# Index 속도 비교
이번에는, index가 있을 때와 없을 때의 쿼리 속도를 비교해보고자 한다.

`order` 테이블에서 index가 있을 때와 없을 때의 조회 쿼리 속도를 비교해 보았다. 실행한 쿼리는 다음과 같다.

```sql
EXPLAIN ANALYZE
SELECT * FROM public.order
WHERE sales_id = 1175;
```

<img width="1086" alt="image" src="https://github.com/yumin00/blog/assets/130362583/444068f0-cbd1-405e-bb3b-c6dfa3f371b3">
<img width="1074" alt="image" src="https://github.com/yumin00/blog/assets/130362583/709e9ad5-9fcd-4d57-9121-4156c26bb888">


첫 번재 사진은 인덱스를 설정하지 않았을 경우이고, 두 번재 사진은 인덱스를 설정한 경우이다.

### 인덱스 설정 X
- Sequential Scan
- 실행시간 : 1.433ms
- 비용: 0.00 to 197.57
- order 테이블의 모든 행을 처음부터 끝까지 검색

### 인덱스 설정 O
- Index Scan
- 0.174ms
- 0.29 to 8.30
- 인덱스 트리 구조를 사용하여 특정 sales_id 값에 직접 접근

인덱스가 설정되어 있을 때와 없을 때의 쿼리 속도와 비용 차이를 명확하게 확인할 수 있다.
인덱스가 설정되어 있다면, PostgreSQL은 `sales_id` 를 기준으로 데이터를 훨씬 더 빠르고 효율적으로 찾아낸다는 것을 알 수 있다!

그렇다면, JOIN문에서도 index 가 효율적으로 사용될까?

`sales`와 `order` 테이블이 존재하고, `order` 테이블에서는 `sales` 테이블의 id인 `sales_id` 가 인덱스로 설정되어 있다.
`order` 테이블에서 `sales_id`를 인덱스로 설정했을 때와 아닐 때의 쿼리 속도를 비교해보자!
실행하고자 하는 쿼리는 다음과 같다.

```sql
EXPLAIN ANALYZE
SELECT o.id, o.order_date, o.status, s.sales_date, s.units_sold
FROM order o
JOIN sales s ON o.sales_id = s.id;
```
<img width="987" alt="image" src="https://github.com/yumin00/blog/assets/130362583/d5bf99e4-8cfb-4d21-bc4d-4e1003dc0021">
<img width="996" alt="image" src="https://github.com/yumin00/blog/assets/130362583/d11a47e4-5811-4471-b222-ea07d034725d">

첫 번째 사진은 인덱스를 설정하지 않았을 경우이고, 두 번째 사진은 인덱스를 설정했을 경우이다.

인덱스를 설정했을 때, 실행 시간이 약간 개선된 것으로 보이지만, 여기서 중요한 점은 실행 시, Seq Scan과 Hash Join이 사용된다는 것이다.
인덱스를 설정했을 경우에도, order 테이블에서 인덱스를 통핸 스캔이 아니라 순차 스캔을 사용하고 있으며, 인덱스가 스캔에 전혀 영향을 주고 있지 않다. 이것의 의미는 쿼리문에서 인데스가 조인 최적화에 활용하지 않고 있다는 것을 의미하기도 하다.
나의 예상으로는, `sales` 와 `order` 테이블은 `sales.id`와 `order.sales_id`가 서로 1:1로 매칭되어 있기 때문에 WHERE 절을 통해 데이터를 찾는 것이 아니라면 인덱스가 효율성을 높이는 역할을 하는 것은 아닌 것 같다.

그러면, WHERE 절을 추가하여 인덱스의 쿼리 속도를 비교해보자.

다음과 같은 쿼리문을 사용해보았다.

```sql
EXPLAIN ANALYZE
SELECT o.id, o.order_date, o.status, s.sales_date, s.units_sold
FROM public.order o
JOIN sales s ON o.sales_id = s.id
WHERE o.sales_id = 1175;
```
<img width="1038" alt="image" src="https://github.com/yumin00/blog/assets/130362583/658e2dcf-0834-4705-8a04-eb315123c180">
<img width="1106" alt="image" src="https://github.com/yumin00/blog/assets/130362583/d9afe087-7353-4eb9-bf2d-28b2f54bfb1b">

첫 번재 사진은 인덱스를 설정하지 않았을 경우이고, 두 번재 사진은 인덱스를 설정한 경우이다.

### 인덱스 설정 X
- Nested Loop Join
- order 테이블에서 Sequential Scan 을 사용하여 조회
- sales 테이블에서 id 에 대한 Index Scan 을 사용하여 조회 (sales.id는 기본키이다)


### 인덱스 설정 O
- Nested Loop Join
- order/sales 모두 Index Scan 사용

인덱스를 설정했을 때, 쿼리의 비용과 속도가 감소하는 것을 볼 수 있다. 이는 order 테이블에서 sales_id를 찾을 때 Index Scan 을 사용함으로써 데이터를 효율적으로 조회했음을 알 수 있다!