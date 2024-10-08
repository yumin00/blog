---
title: "Spark에 대해 알아보자"
date: 2024-05-30T21:16:10+09:00
draft: false
categories :
- DataEngineering
---

(해당 글에는 패스트캠퍼스의 한 번에 끝내는 데이터 엔지니어링 초격차 패키지 강의의 내용이 포함되어 있습니다.)

## 1. Spark 는 왜 생겨났을까?
### MapReduce 한계 극복
이전 [하둡](https://yumin.dev/p/hadoop%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90/)에서 공부한 것을 떠올려보자. 하둡은 HDFS를 통해 데이터를 분산 저장하고 맵리듀스를 통해 데이터를 병렬 처리하며 동작한다.

하둡의 HDFS는 디스크 I/O로 동작하기 때문에 실시간 데이터를 처리하기에는 속도 측면에서 부적합하다. 하지만 실시간 데이터 처리에 대한 니즈가 발생하면서 생겨난 것이 바로 Apache Spark 이다.

또한, 하둡의 MapReduce는 자원의 할당과 해제가 빈번하게 이루어지기 때문에 성능의 저하를 가져온다. 또한, 맵리듀스를 진행하면서 중간 결과 파일을 하둡에 올려놓기 때문에 지연시간이 발생하고,
실제로 최종 결과물이 아닌 중간 결과를 하둡에 저장하는 것 자체가 부하 문제를 가져올 수 있다.

Spark는 RDD 데이터 모델과 인메모리 연산을 지원하여 하둡보다 최소 100배 이상 빠르게 작동하며 이는 MapReduce의 한계를 극복할 수 있게 해준다.(실제로 그러한지에 대해서는 나중에 직접 테스트를 해보고자 한다.)

그럼 지금부터 Spark에 대해 더 자세히 알아보자.

## 2. Spark
Spark는 대규모 데이터 처리와 분석을 위한 오픈 소스 분산 컴퓨팅 시스템이다.

### API로 정형화된 데이터 처리 모델
MapReduce는 데이터 변환 > 데이터 필터링 > 데이터 재배치 > 데이터 집계 function을 통해 이루어진다. 스파크는 맵리듀스의 방식을 API로 정형화화여 데이터 처리 작업을 직접 생성하는 것보다 더 쉽게 코드를 작성할 수 있게 해준다.

### 분산 컴퓨팅 시스템
분산 컴퓨팅 시스템은 여러 대의 컴퓨터가 협력하여 하나의 큰 작업을 수행하는 시스템을 의미한다. 즉, 하나의 컴퓨터가 아닌 여러 컴퓨터 자원을 사용할 수 있다는 것이다.
이 시스템은 작업을 병렬로 나누어 수행함으로써 성능을 향상시키고, 대규모 데이터 처리 및 복잡한 계산을 보다 효율적으로 할 수 있게 한다.

대량의 데이터를 여러 컴퓨터 자원을 사용하여 처리함에도 신뢰성 있게 처리할 수 있다는 특징이 있다.

### 인메모리 연산
Spark의 가장 큰 특징은 인메모리 연산이다. 디스크가 아닌 메모리 내에서 연산을 지원하기 때문에 빠른 속도를 제공한다. 이러한 인메모리 연산을 사용하여 하둡 생태계를 보완하는 기술로 자리잡게 되었다.

### RDD (Resilient Distributed Datasets)
RDD란, Spark의 핵심 데이터 구조로 대규모 데이터 처리를 위한 추상화 계층이다. RDD는 Spark의 분산 데이터 처리를 효율적으로 수행할 수 있도록 설계되어 있다.

RDD는 다음과 같은 특징을 갖는다.

- 불변성: RDD는 한번 생성되면 변경할 수 없다. 데이터를 변환하면 새로운 RDD를 생성한다. 이는 병렬 처리를 안전하게 만들고, 데이터의 일관성을 보장한다.
- 분산성: RDD는 클러스터의 여러 노드에 걸쳐 분산 저장된다. 이를 통해 대규모 데이터를 효율적으로 처리할 수 있다.

### 스트리밍 데이터 처리
Spark Streaming을 통해 실시간 데이터 스트리밍 처리를 할 수 있다.

### 배치 처리
Batch 처리 기능을 제공하기 때문에 배치 처리와 스트리밍 처리를 통해서 통합 개발할 수 있다.

## 4. Spark 구성 요소
구성 요소에 대해 알아보기 전, 구성 요소를 공부하면 자주 나오는 개념들에 대해 먼저 간단히 알아보고 가자.

- 클러스터(Cluster): 여러 대의 노드(서버)로 구성된 환경
- 애플리케이션(Application): [사용자 코드] Spark에서 실행될 작업(job)과 그 작업을 구성하는 여러 단계(stage)를 포함
- SparkContext: 클러스터와 통신을 담당하는 객체
- 작업(Job): 데이터를 처리하는 단위

![image](https://github.com/yumin00/blog/assets/130362583/46601663-ba1b-48be-9d36-95bcafa7bfa6)

### 드라이버 프로그램 (Driver Program)
Spark가 어플리케이션으로서 기능을 수행하기 위해서는 SparkContext가 시작되어야 한다. 드라이버 프로그램은 SparkContext를 시작시키는 역할을 한다.

즉, 드라이버 프로그램은 스파크의 중앙 처리 장치이다. SparkContext 를 시작하고, 사용자가 작성한 애플리케이션 코드를 실행하는 역할을 담당한다.

드라이버 프로그램이 어플리케이션을 실행시킬 때, 코드 내용을 보고 job의 순서를 나눈다(DAG - Logical Plan). 각 job은 stage로 나누고, 각 stage(수행할 논리적인 공간)는 작은 task(실제 실행 단위)로 나뉜다. 그리고 Task는 Task Scheduler에게 전달된다.

task로 나눈 다음, 드라이버 프로그램은 클러스터 매니저에게 필요한 excutor 개수와 실행할 테스크를 알려주며 자원 할당을 요청한다.  

드라이버 프로그램은 마지막에 각 task의 중간/최종 결과를 수집하고, 최종 결과를 사용자에게 반환한다.
 
### 클러스터 매니저 (Cluster Manager)
클러스터 매니저는 클러스터의 자원을 관리하며 드라이버 프로그램이 자원 할당을 요청하면 이에 응답하는 역할을 한다.

실제로 리소스를 할당하여 실행해야 할 때, 클러스터 매니저가 가지고 있는 워커 노드에게 작업을 실행해야 한다는 것을 알리고, 실제로 워커 노드에서 사용하는 CPU, Memory의 성능/상태 정보를 관리한다.

워커 노드에서 테스크가 완료되면, 최종 실행 정보도 함께 트래킹한다.

만약, 워커 노드가 죽게 되면 해당 워커 노드에 실행하고 있던 테스크도 중지되게 되는데 클러스터 매니저는 해당 테스크의 정보를 유지한 다음, 해당 테스크를 다른 워커 노드에게 재스케줄하는 역할도 함께 한다.

클러스터 매니저에는 다양한 종류가 있다.

- Standalone Cluster Manager: 기본 클러스터 매니저로, 독립적인 Spark 클러스터를 설정할 때 사용된다.
- Apache Mesos: 복잡한 클러스터 환경에서 자원 관리를 위해 사용된다.
- Hadoop YARN: Hadoop 생태계와 통합된 클러스터 매니저로, Hadoop 클러스터에서 Spark를 실행할 때 사용된다.
- Kubernetes: 컨테이너화된 환경에서 클러스터 자원을 관리할 때 사용된다.

### 실행기(Executor)
** Spark를 Submit할 때 사용자는 executor의 개수, executor의 memory, executor의 core 수를 지정할 수 있다.

실행기는 각 노드에서 실행되는 프로세스로, 실제로 작업을 수행한다. 실제로 실행되는 테스크들의 데이터는 모두 실행기의 Disk에 저장된다.

- 작업 실행: 드라이버 프로그램에서 할당된 작업을 실행한다.
- 데이터 저장: 메모리 또는 디스크에 데이터를 저장한다.
- 작업 결과 반환: 작업 결과를 드라이버 프로그램에 반환하다.

### RDD(Resilient Distributed Dataset)
RDD는 분산된 데이터 세트를 표현한다. RDD는 한 번 생성되면 변경할 수 없고 여러 노드에 분산되어 저장된다. 장애가 발생하면 RDD는 자동으로 복구할 수 있다는 특징을 가진다.

## 5. Spark 동작 방식
실제로, 어플리케이션 코드가 실행됐을 때 Spark 가 어떻게 동작하는지에 대해 더 자세히 살펴보자.

### [1] 드라이버 프로그램: SparkContext 생성
드라이버 프로그램은 SparkContext를 생성하고, 이로 인해 클러스터 매니저와 연결된다.

### [2] 드라이버 프로그램: DAG 생성 및 스테이지 분할
![image](https://github.com/yumin00/blog/assets/130362583/18497031-dc24-4f69-aada-7ab0efc4994a)

드라이버 프로그램은 사용자가 작성한 트랜스포메이션(ex. `map`, `filter`, `flatMap`)을 바탕으로 DAG를 생성하고 이를 여러 스테이지로 분할한다. 이때 DAG는 각 스테이지가 실행되어야 하는 순서를 정의한다.

### [3] 드라이버 프로그램: 테스크 생성
드라이버 프로그램은 스테이지를 다시 테스크로 분할하고, 클러스터 매니저에게 해당 테스크를 실행할 실행기를 요청한다.

### [4] 클러스터 매니저: 테스크 실행
클러스터 매니저는 리소스를 할당하여 테스크를 실행시킨다.

### [5] 실행기: 결과 저장 및 전송
각 실행기는 테스크의 실행 결과를 메모리에 저장하고, 결과를 직렬화하여 드라이버 프로그램에게 전송한다.

### [6] 드라이버 프로그램: 결과 반환
드라이버 프로그램은 직렬화된 결과 데이터를 역직렬화하여 원래 데이터 형태로 변환한다. 드라이버 프로그램은 각 실행기로부터 받은 부분 결과를 수집하여 하나의 최종 결과로 수집하여 사용자에게 반환한다.

## Spark 사용 시 주의할 점
### Driver Program
#### Client-Mode
드라이버 프로그램은 하나이고, executor는 여러 개일 수 있다. executor들의 결과값은 모두 드라이버 프로그램으로 모이기 때문에 결과값의 데이터가 클 경우, 드라이버 프로그램이 오버헤드(Out Of Memory)로 죽는 경우가 발생할 수 있다.

이때, 드라이버 프로그램을 자동으로 살릴 수 있는 방법은 없기 때문에 드라이버 프로세스를 모니터링하고 드라이버에 문제가 발생할 경우 대처할 방법을 미리 강구해야 한다.

#### Cluster-Mode
드라이버가 있는 컨테이너가 죽으면 기본적으로 드라이버는 다시 시작한다. 하지만 이때 executor들이 전달한 데이터가 유지된다는 보장은 없기 때문에 재시작에 대한 고려가 필요하다.

### Executor
executor의 수를 늘릴수록 동시에 처리할 수 있는 데이터가 늘어난다. executor 수가 증가하면 병렬연산 속도가 빨라지는 것을 기대할 수 있다.

하지만 데이터 양에 비해 너무 많은 executor를 사용하게 되면 불필요하게 네트워크를 많이 사용하고 이동하는 데이터가 많아지기 때문에 오히려 수행 속도가 느려지고 시스템 부하가 높아질 수 있다.

![image](https://github.com/user-attachments/assets/c8c6a19e-415d-4554-9fb2-730867bdb770)


따라서 프로그램의 로직 / 데이터의 사이즈 / 리소스의 수 / 기대하는 퍼포먼스 를 고려하여 executor를 설정해야 한다.

### Idempotent
executor는 클러스터 매니저와 드라이버에 의해 다시 시작될 수 있다. 따라서, job은 언제든지 다시 시작될 수 있다는 것을 가정하고 작성하는 것이 가장 좋다. 하나의 작업을 멱등성 있게 작성해야 실제 데이터가 잚못 처리되는 것을 막을 수 있다.




## 5. Hadoop 내 에서의 Spark
Spark는 Hadoop과 통합되어, Hadoop EcoSystem 내에서 강력한 데이터 처리 기능을 제공한다. 특히, HDFS와 YARN 과의 통합이 대표적인데 이에 대해 알아보자.

### Spark와 HDFS의 통합
Spark는 HDFS에 저장된 데이터를 직접 읽고 쓸 수 있다. 이를 통해 대규모 데이터를 분산 처리할 수 있다.

### YARN 클러스터 모드에서 Spark 실행
Spark는 YARN 클러스터 매니저를 통해 자원을 할당받아 작업을 실행할 수 있다. 이를 통해 하둡 클러스터의 자원을 효율적으로 사용할 수 있다.

### Spark Streaming
Spark Streaming을 사용하면 실시간 데이터 스트림을 처리할 수 있다.