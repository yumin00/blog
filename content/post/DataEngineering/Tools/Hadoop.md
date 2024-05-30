---
title: "Hadoop에 대해 알아보자"
date: 2024-05-20T20:18:15+09:00
draft: false
categories :
- DataEngineering
---

# 1. Hadoop은 왜 생겨났을까?
하둡이 무엇인지 알아보기 전에, 먼저 하둡이 왜 탄생했는지에 대해 먼저 알아보자.

## 대용량 데이터 처리의 필요성
1990년대 후반과 200년대 초반, 기술의 반전으로 인해 데이터의 양은 급격하게 증가하기 시작했다.
점점 많아지는 데이터를 기존 데이터베이스인 RDBMS로 처리하기에는 한계가 있었다.

## 구글의 논문 발표
2003년과 2004년에 구글은 대용량 데이터 처리 문제를 해결하기 위해 두 가지 핵심 논문을 발표했다.

- 구글 파일 시스템(Google File System, GFS): 대규모 데이터를 분산된 방식으로 저장하고 관리할 수 있는 파일 시스템
- 맵리듀스(MapReduce): 대규모 데이터를 병렬로 처리할 수 있는 프로그래밍 모델

이 논문들은 대용량 데이터 처리를 효율적으로 수행할 수 있는 방법론을 제시했으며, 이는 이후에 하둡의 핵심 아이디어가 되었다.

## 하둡의 탄생
더그 커팅(Doug Cutting)과 마이크 카파렐라(Mike Cafarella)는 구글의 논문에서 영감을 받아 2006년에 하둡을 개발하기 시작했다.
더그 커팅은 원래 루씬(Lucene)이라는 오픈 소스 검색 엔진 프로젝트를 진행 중이었는데, 빅데이터 처리에 대한 필요성을 느끼고 하둡 프로젝트로 전환하게 되었다.
하둡은 초기에는 야후(Yahoo)에서 사용되었고, 이후 아파치 소프트웨어 재단의 프로젝트로 발전했다.

## 하둡의 목적 및 특징
하둡은 다음과 같은 특징을 가지고 대용량 데이터를 효율적으로 처리하고 저장하기 위해서 생겨났다.

- 대규모 데이터 처리: 기존 RDBMS가 처리할 수 없는 대용량 데이터를 효율적으로 처리
- 분산 저장: 데이터를 여러 서버에 분산 저장하여 데이터 손실을 방지
- 병렬 처리: 데이터를 병렬로 처리하여 처리 속도를 크게 향상
- 저비용: 상용 하드웨어를 활용하여 효율적인 비용으로 대규모 데이터 처리 시스템 구축

그럼 지금부터 하둡은 어떤식으로 동작하길래 위와 같은 특징을 가지고 대규모 데이터 처리를 할 수 있는지에 대해 더 자세히 알아보자!

# 2. Hadoop
하둡은 어떻게 대규모 데이터를 처리할 수 있는 것일까? 하둡은 여러 가지 요소들을 사용하여 대용량 데이터를 효율적으로 처리한다.

## (1) 분산 파일 시스템 (HDFS - Hadoop Distributed File System)
HDFS는 데이터를 여러 서버(노드)에 분산 저장한다.

### 블록 단위 저장
HDFS는 대규모 데이터를 여러 블록으로 나누어 여러 노드에 저장한다.

> 블록
> 
> 블록이란, HDFS에서 데이터를 저장할 때 사용하는 고정 크기의 데이터 조각이다. 기본 블록 크기는 128MB이지만 설정에 따라 조정할 수 있다.

#### 블록 관리
![image](https://github.com/yumin00/blog/assets/130362583/901bb69a-e7a2-4866-9359-66fdffb3d939)

- 네임노드 (NameNode)
  - 파일 시스템의 메타데이터를 관리한다.
  - 파일과 디렉터리의 구조, 각 블록의 위치 정보를 유지한다.
  - 블록 복제본의 위치를 추적하고, 데이터노드의 상태를 모니터링한다.
  - 데이터노드 장애 시 자동으로 블록을 재복제하여 데이터의 가용성을 유지한다.
- 데이터노드 (DataNode)
  - 실제 데이터를 저장하는 노드
  - 주기적으로 네임노드에 하트비트를 보내 블록 상태를 보고한다.

    
#### 블록은 왜 사용할까?
1. 단순화된 메타데이터 관리

파일 시스템에서 각 파일의 메타데이터를 관리하는 것은 많은 오버헤드를 발생시킨다. 블록 단위로 데이터를 저장함으로써, HDFS는 파일 메타데이터를 간소화하고, 네임노드가 관리해야 하는 메타데이터의 양을 줄인다.

2. 데이터 복제

HDFS는 각 블록을 기본적으로 세 개의 복제본(replica)을 유지하고, 각 복제본들은 서로 다른 데이터노드(Data Node)에 저장된다. 이를 통해 하나의 노드가 고장나더라도, 저장된 복제본을 이용하여 데이터 손실을 방지한다.

여기서 나는 의문점이 생겼다. 하둡은 대용량 데이터를 효율적으로 처리하기 위함인데, 복제본을 만들면 너무 많은 용량을 차지하는 것은 아닐까?

하지만 하둡은 기본적으로 저비용의 상용 하드웨어를 사용하기 떄문에, 저장 공간 비용이 고성능 전용 하드웨어에 비해 상대적으로 낮다고 한다. 또한 아래서 말하는 다른 요인들이 저장 공간에 대한 추가 비용을 상쇄한기 때문에 이러한 정책을 채택한 것 같다.

3. 병렬 처리

데이터를 블록 단위로 나눠 분산 저장하고 처리함으로써 여러 노드에서 동시에 읽고, 쓰기를 할 수 있어 데이터 처리 속도를 향상시킬 수 있다.


4. 데이터 현지성 (Data Locality)

HDFS는 작업을 수행할 때 데이터가 저장된 노드에서 직접 처리하도록 하여, 네트워크 트래픽을 줄이고 성능을 최적화한다.

5. 유연한 확장

블록 단위로 데이터를 저장함으로써, 새로운 데이터노드를 추가하는 것만으로도 쉽게 저장 용량과 처리 능력을 확장시킬 수 있다.

### 파일 저장 플로우
![image](https://github.com/yumin00/blog/assets/130362583/b906f2cd-3aa9-443b-ae0d-129331760df2)
그러면 실제로 데이터가 들어왔을 때, HDFS는 어떻게 작동하는지에 대해 알아보자.

1. 파일 분할

클라이언트가 파일을 업로드하려고 하면, HDFS는 이를 블록 단위로 분할한다.

2. 네임노드에 메타데이터 등록

네임노드에 해당 파일의 메타데이터를 등록한다.

3. 네임노드의 블록 지정

네임노드의 각 블록을 저장할 데이터노드를 지정하여 목록을 반환한다.

4. 데이터노드에 블록 저장

클라이언트는 네임노드로부터 지정받은 데이터노드에 해당 블록을 저장한다. 그리고 각 데이터노드들은 저장이 완료되면 다음 데이터노드에게 해당 블록을 전달한다. 

5. 데이터 저장 및 확인

각 데이터노드는 블록들을 자신의 디스크에 저장하고, 블록이 성공적으로 저장되면 전달받은 데이터노드에게 확인 메시지를 전송한다. 블록 복제본이 모두 저장되면 첫 번째 데이터노드는 이를 클라이언트에게 전달한다.

6. 네임노드의 메타데이터 업데이트

클라이언트는 네임노드로부터 모든 블록이 성공적으로 저장되었음을 확인 받으면, 파일 업로드가 완료되고 네임노드는 파일 메타데이터를 최종적으로 업데이트한다.

### 파일 읽기 플로우
![image](https://github.com/yumin00/blog/assets/130362583/65813f04-6631-486f-b55d-5c1f0ca2ef9c)
1. 파일 읽기를 요청

클라이언트가 HDFS에 파일 읽기를 요청한다.

2. 네임 노드의 메타데이터 확인

네임 노드는 요청된 파일이 어떤 블록에 저장되어있는지 메타데이터를 확인하여 파일이 저장된 블록 리스트를 반환한다.

3. 데이터노드에서 데이터 읽기

클라이언트는 블록이 저장된 각 데이터노드에 접근하여 블록 조회를 요청한다. (병렬 혹은 순차적으로)

4. 데이터노드의 응답

데이터노드는 요청된 블록을 전송한다.


## (2) 병렬 처리 모델 (MapReduce)
HDFS를 통해 데이터를 분산 저장하게 되면, 문제는 분산되어 있는 데이터들을 한 번에 처리할 수 없다. 이러한 문제를 해결해주는 것이 바로 맵리듀스이다.

맵리듀스란, 데이터를 병렬으로 처리할 수 있는 프로그래밍 모델이다.
맵리듀스는 맵(Map) 단계와 리듀스(Reduce) 단계로 나뉘며, 이를 통해 데이터를 병렬로 처리하고 집계한다.

맵리듀스를 통해 대규모 데이터를 효율적으로 분석하고 집계할 수 있다. 맵리듀스는 데이터를 분산 처리하여 작업 속도를 높이고, 데이터 처리의 신뢰성을 보장하는 도구이다.

### 맵 단계
맵은 데이터를 수직화하여 그 데이터를 각각 종류별로 모으는 단계이다.

맵 단계는 데이터를 작은 청크로 나누어서 특정 키-값 쌍으로 변환하여 키를 기준으로 데이터를 정렬화하고 그룹화한다.

### 리듀스 단계
리듀스는 필터링과 sorting을 거쳐 데이터를 뽑아내는 단계이다.

리듀스 단게에서는 맵 단계에서 발생한 데이터들의 중복을 제거하고 원하는 데이터를 추출한다.

#### 요약
1. 디스크로부터 데이터를 읽어온다.
2. Map 단계에서 데이터를 키-값 으로 정렬화/그룹화한다.
3. Reduce 단계에서 중복된 키를 가진 데이터를 제거하고, 원하는 데이터를 추출한다.
4. 결과를 디스크에 저장한다.

### 예시
예를 들어서, 단어 빈도수를 계산하고 싶은 상황이라고 가정해보자. 맵 단계에서는 파일의 읽어 각 단어를 키로, 빈도를 값으로 출력한다. 리듀스 단계에서는 같은 단어를 키로 가진 모든 빈도를 합산하여 최종 빈도를 출력해준다.

이렇게 맵리듀스를 사용함으로써 대용량 데이터를 효율적으로 분석하고 집계할 수 있는 것이다.

### 구성
![image](https://github.com/yumin00/blog/assets/130362583/d8f589b4-2d64-49e6-a674-098c24526c4a)
하둡에서 맵리듀스는 클라이언트, 잡 트래커(Job Tracker), 태스크 트래커(Task Tracker)로 구성된다.
- 클라이언트 : 분석하고자 하는 데이터를 잡(JOB)의 형태로 JobTracker 에게 전달
- 잡 트래커 : 네임노드에 위치, 하둡 클러스터에 등록된 전체 잡(JOB) 을 스케줄링하고 모니터링
- 태스크 트래커: 데이터 노드에서 실행되는 데몬, Task를 수행하고, 잡 트래커에게 상황을 보고

## (3) 하둡 에코 시스템 (Hadoop EcoSystem)
하둡 에코 시스템은 하둡 프레임워크를 중심으로 다양한 도구와 기술들이 결합되어 빅데이터 처리 및 분석을 가능하게 하는 생태계를 의미한다. 하둡 에코 시스템에 HDFS와 맵리듀스가 포함되는 것이며 그 외에도 다양한 요소들이 존재한다.

### YARN (Yet Another Resource Negotiator)
> 클러스터 자원
> 
> 하둡 클러스터 내에서 데이터를 처리하고 저장힉 위해 사용되는 모든 하드웨어 및 소프트웨어 자원

YARN은 클러스터 자원을 관리하고 어플리케이션의 실행을 조정하는 자원 관리 시스템이다. 다양한 데이터 처리 엔진(ex. 맵리듀스)을 지원하여 자원 할당과 작업 스케줄링을 수행한다.

### HBase
HBase는 구글 BigTable을 기반으로 개발된 비관계형 데이터베이스이다. Hadoop과 HDFS 위에 BigTable과 같은 기능을 제공해준다.
네이버 라인 메신저에서는 HBase를 적용한 시스템 아키텍쳐를 발표했다고도 하는데, 이에 대해서는 좀 더 자세히 공부해보고자 한다.


## 하둡의 장단점
### 장점
- 오픈소스이기 때문에 비용 부담이 적다.
- 일부 장비에 장애가 있더라도 전체 시스템 사용에 영향이 적다.

### 단점
- HDFS에 있는 데이터를 변경할 수 없다.
- 실시간 데이터 분석과 같은 스트리밍 서비스에는 부적합하다.

## 하둡을 직접 설치해 보자.
### 1. 하둡 설치
먼저, 하둡을 설치하자.
```
brew install hadoop
```

### 2. 환경 변수 수정
이제 설치한 하둡 위치로 이동해보자. 1번 명령어를 통해 얻은 경로로 2번 명령어를 통해 이동해보자.
그리고 환경 변수가 있는 파일들을 수정하기 위해 3번 명령어로 이동해야한다.

```
## 1. hadoop 경로 확인
brew info hadoop

## 2. hadoop 경로로 이동
cd /opt/homebrew/Cellar/hadoop/3.4.0

## 3. 파일 수정을 위해 이동
cd libexec/etc/hadoop
```

#### 2-1. Java 버전 명시
Hadoop은 Java로 작성되었기 때문에, Java 런타임 환경이 반드시 필요하다. 때문에 `hadoop_env.sh`에 Hadoop이 실행할 때 사용할 JAVA 버전을 명시적으로 지정해줘야 한다.

```shell
/usr/libexec/java_home
```

위 명령어를 통해 Java의 설치 경로를 알 수 있다.

```shell
open hadoop-env.sh
```

파일을 열고, 맨 아래에 자바 경로를 표시해주면 된다.

```shell
export JAVA_HOME="/Library/Internet Plug-Ins/JavaAppletPlugin.plugin/Contents/Home"
```

#### 2-2. HDFS 주소 설정
여러 구성 파일 중에서 core-site.xml은 Hadoop의 기본 설정을 담고 있다. core-site.xml에 URI를 설정하여 해당 주소를 파일 시스템으로 사용할 수 있도록 설정해주어야 한다. 

```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0/libexec/etc/hadoop
vi core-site.xml
```

core-site.xml을 열어서 다음 설정을 추가해줘야 한다.

```shell
<configuration>
  <property>
    <name>fs.defaultFS</name>
    <value>hdfs://localhost:9000</value>
  </property>
</configuration>
```

- `fs.defaultFS`: "기본 파일 시스템 주소" 를 의미한다.
- `hdfs://localhost:9000`: HDFS는 localhost:9000 주소를 사용하라는 뜻이다.

#### 2-3. HDFS 옵션 설정
다음으로, HDFS에서 파일의 복제본 개수를 설정해보자.

```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0/libexec/etc/hadoop
vi hdfs-site.xml
```

hdfs-site.xml 을 열어서 다음 설정을 추가해보자.

```shell
<configuration>
  <property>
    <name>dfs.replication</name>
    <value>1</value>
  </property>
  <property>
    <name>dfs.datanode.data.dir</name>
    <value>/usr/local/hadoop/data</value>
  </property>
</configuration>
```

- `dfs.replication`: HDFS에서 파일의 복제본(복제 블록)의 개수를 의미한다.
- 3: 각 파일의 복제본을 3개를 만들겠다는 뜻이다.

#### 2-4. MapReduce 설정
다음으로는 MapReduce 작업이 올바르게 실행될 수 있도록 필요한 환경과 경로를 지정해보자.

```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0/libexec/etc/hadoop
vi mapred-site.xml
```

mapred-site.xml 을 열어서 다음 설정을 추가해보자.

```shell
<configuration>
  <property>
    <name>mapreduce.framework.name</name>
    <value>yarn</value>
  </property>
  <property>
    <name>mapreduce.application.classpath</name>
    <value
      >$HADOOP_MAPRED_HOME/share/hadoop/mapreduce/*:$HADOOP_MAPRED_HOME/share/hadoop/mapreduce/lib/*</value
    >
  </property>
</configuration>
```

- mapreduce.framework.name / yarn: MapReduce 프레임워크를 YARN(Yet Another Resource Negotiator) 를 사용하겠다는 의미이다.

> YARN
> 
> 하둡의 자원 관리와 스케줄링을 담당하는 프레임워크로, 맵리듀스뿐만 아니라 다양한 분산 컴퓨팅 작업을 실행할 수 있게 도와준다. yarn으로 설정하면, 모든 맵리듀스 작업이 YARN 클러스터 위에서 실행된다.

#### 2-5. YARN 환경 설정
```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0/libexec/etc/hadoop
vi yarn-site.xml
```

```shell
<configuration>
  <property>
    <name>yarn.nodemanager.aux-services</name>
    <value>mapreduce_shuffle</value>
  </property>
  <property>
    <name>yarn.nodemanager.env-whitelist</name>
    <value
      >JAVA_HOME,HADOOP_COMMON_HOME,HADOOP_HDFS_HOME,HADOOP_CONF_DIR,CLASSPATH_PREPEND_DISTCACHE,HADOOP_YARN_HOME,HADOOP_HOME,PATH,LANG,TZ,HADOOP_MAPRED_HOME</value
    >
  </property>
</configuration>
```

### 3. 하둡 실행
#### 3-1. 파일 시스템 포맷
```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0
hdfs namenode -format
```
기존에 있는 데이터를 모두 삭제하고, 새 파일 시스템 구조를 만드는 작업이다. 네임 노드를 초기화하고, 파일 시스템의 메타데이터를 생성하는 작업이다.

#### 3-2. 사용자 디렉토리 생성
```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0
bin/hdfs dfs -mkdir /user
bin/hdfs dfs -mkdir /user/<username>
```
사용자별 디렉토리를 만들면, 사용자가 자신의 데이터를 독립적으로 관리할 수 있기 때문에 위와 같이 별도의 디렉토리를 만드는 것이다.

#### 3-3. 하둡 실행
```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0
sbin/start-all.sh
```
해당 명령어를 통해 하둡을 실행시킬 수 있다. 

#### 3-4. 실행 확인
```shell
## 위치: /opt/homebrew/Cellar/hadoop/3.4.0
jps
```
해당 명령어를 통해 하둡이 정상적으로 실행되고 있는 것을 확인할 수 있다!
```shell
15042 Main
53250 Jps 
627 Main
50055 ResourceManager
50152 NodeManager
633 
52043 NameNode
52284 SecondaryNameNode
```

각 행은 다음과 같은 정보를 제공한다.
- 프로세스 ID(PID): 첫 번째 숫자는 각 프로세스의 고유한 식별자이다.
- 프로세스 이름: 두 번째 필드는 프로세스의 이름이다. 이 이름은 실행 중인 Java 클래스나 애플리케이션을 나타낸다.

해당 명령어를 통해 클러스터의 상태를 모니터링할 수 있고 문제를 진단해볼 수 있다.

localhost 로 접속해서 직접 확인해 볼 수 있다.

- Cluster status : http://localhost:8088
- HDFS status : http://localhost:9870
- Secondary NameNode status : http://localhost:9868

## Hadoop의 발전
하둡은 빅데이터를 더 효율적으로 저장하기 위해서 탄생했다. 하지만 데이터 프로세싱 속도가 느리게 되어 SPARK 가 탄생하게 되었다.(SPARK는 많은 자원이 필요한 대신, 맵리듀스보다 속도가 더 빠르다고 한다) 그리고 하둡 내에서 직접 데이터 분석을 하기 위해 HIVE도 탄생하게 되었다.

다음에는 하둡을 통해 탄생하게 된 Tool들에 대해 공부해보고자 한다.