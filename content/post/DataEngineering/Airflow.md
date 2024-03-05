---
title: "Airflow 에 대해 알아보자"
date: 2024-02-27T19:53:19+09:00
draft: false
categories :
- DataEngineering
---

# Airflow 에 대해 알아보자
## Airflow
Apache Airflow는 초기 에어비엔비(Airfbnb) 엔지니어링 팀에서 개발한 워크플로우 오픈 소스 플랫폼이다.
> 워크플로우란?
> 
> 의존성으로 연결된 작업(Task)들의 집합

파이썬 코드로 워크플로우를 작성하고 스케줄링, 모니터링할 수 있다.

### DAG
하나의 워크플로우를 DAG 라고 칭한다. DAG 는 Directed Acyclic Graph 의 약자로 방향성을 가졌지만 순환하지 않는 그래프를 말한다.

DAG 안에는 하나 이상의 Task가 포함되고 각 Task들의 순서와 의존성을 정의할 수 있다.

```python
t1 >> t2 >> t3
```
![image](https://github.com/yumin00/blog/assets/130362583/f81fde66-8655-4085-9de7-f7baea618919)

```python
t1 >> t4
t1 >> t2 >> t3
```
![image](https://github.com/yumin00/blog/assets/130362583/4847b279-aab5-458d-9b36-4b370cdcb45b)

### Operator
Task를 정의할 때 Operator를 사용한다. Operator Type은 다음과 같다.

- Action Operators
  - 기능, 명령을 수행하는 오퍼레이터
  - 실제 연산을 수행하는 오퍼레이터
  - 내장 Operators: BashOperator, PythonOperator, HttpOperator
    - BashOperator: bash command 실행
    - PythonOperator: 파이썬 함수 실행
    - HttpOperator: http request 실행
    - (더 많은 오퍼레이터는 [공식문서](https://airflow.apache.org/docs/) 참고할 수 있다)
- Transfer Operator
  - 특정 시스템에서 다른 시스템으로 데이터를 이동시키는 오퍼레이터
- Sensor Operator
  - 조건이 만족할 때까지 기다렸다가, 조건이 충족되면 다음 Task 실행
  - ex) Kafka Topic에 이벤트가 consume 되면 다음 Task 실행

## Airflow Architecture
![image](https://github.com/yumin00/blog/assets/130362583/e7d2eb57-7170-463c-b74a-c86a8124fa87)

에어플로우는 Web Server, Scheduler, Worker, MetaStore, Executor 5개의 기본 구성으로 이루어져있다.
- Web Server
  - 웹 대시보드 UI 제공
  - 스케줄러에서 분석한 DAG 시각화 및 DAG 실행결과 확인할 수 있는 인터페이스 제공
![image](https://github.com/yumin00/blog/assets/130362583/047762bc-da9b-4df5-a3ac-621a6c74ed8d)

- Scheduler
  - DAG 분석 및 DAG의 스케줄이 지난 경우 Worker에 DAG의 Task 예약
- Worker
  - 예약된 Task 실행
- MetaStore
  - DAG, Task 등의 메타 데이터 관리
- Executor
  - Task가 어떻게 실행되는지 정의

## Airflow 동작 원리
1. 유저가 새로운 DAG를 작성
2. Web Server와 Scheduler 는 이를 읽어옴
3. Worker는 Scheduler가 예약한 Task를 실행 및 결과 반환
4. 유저는 웹 인터페이스를 통해 Task 실행과 결과를 모니터링

![image](https://github.com/yumin00/blog/assets/130362583/f3817fe9-64c9-451c-a662-6317bf752553)

![image](https://github.com/yumin00/blog/assets/130362583/a32b114b-c819-496c-a780-936140dd894a)

## Airflow 장단점
### 장점
- 스케줄링 기능으로 DAG에 정의된 특정 시점에 트리거 할 수 있다.
- UI가 제공되기 때문에 쉽게 DAG의 과정 및 결과를 모니터링 할 수 있다.
- 확장 및 다양한 시스템과 통합이 가능하다. (DB, Kafka, Kubernetes)
- 역할 기반 접근 제어를 통해 자신에게 할당된 DAG만 제어가 가능하다.
![image](https://github.com/yumin00/blog/assets/130362583/f03e5e1b-7778-4eba-9da4-8e42b849d6c0)

### 단점
- 아키텍처 이해도가 필요하다.
- 유지보수를 위해 DAG 정책과 Operator 모듈화를 통한 관리가 필요하다.