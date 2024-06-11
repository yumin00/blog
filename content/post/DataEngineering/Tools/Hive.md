---
title: "Hive에 대해 알아보자"
date: 2024-06-11T23:07:14+09:00
draft: true
categories :
- DataEngineering
---

## Hive 탄생 배경
Hive에 대해 알기 전에 먼저 탄생 배경에 대해 알아보자.

Hive는 Facebook에서 개발되었으며, Hadoop을 사용하여 대용량 데이터를 처리하기 위한 도구로 시작되었다.
Facebook은 수십 페타바이트에 달하는 데이터를 관리하고 **분석**할 수 있는 방법이 필요했기 때문에, SQL과 유사한 언어를 통해 데이터 분석을 쉽게 할 수 있도록 Hive를 만들었다.

즉, Hive는 Hadoop 내에서 대용량 데이터를 분석하기 위해 탄생했다.

## Hive
Apache Hive는 대용량 데이터 집합을 위한 데이터 웨어하우스 인프라이다. 

Hive는 HDFS에 저장되어 있는 데이터를 RDB처럼 데이터베이스, 테이블과 같은 형태로 정의하는 방법을 제공하기 때문에, SQL과 유사한 HiveQL 쿼리를 사용할 수 있도록 한다. 이를 통해, HDFS에 있는 데이터를 쿼리하고 분석할 수 있는 기능을 제공한다!

- 테이블 (Tables): Hive에서 데이터는 테이블로 조직된다. 각 테이블은 스키마(구조)를 가지고 있으며, 데이터는 HDFS에 저장된다.
- 데이터베이스 (Databases): 관련된 테이블을 논리적으로 그룹화한 것이다.
- 파티션 (Partitions): 테이블을 더 작은 단위로 나눠 효율적인 데이터 관리와 빠른 쿼리를 가능하게 한다.
- 버킷 (Buckets): 파티션된 데이터를 더 작은 부분으로 나눠 데이터 샘플링 및 효율적인 조인을 가능하게 한다.

## Hive 구성요소
![image](https://github.com/yumin00/blog/assets/130362583/bba5c937-efa1-44c2-9df9-ae9429c2a4bd)

- UI: 사용자가 쿼리 및 작업을 할 수 있는 인터페이스
- Driver: 쿼리를 입력받고, 작업을 처리한다.
- Compiler: MetaStore를 참고하여 쿼리 구문을 분석하고 실행 계획을 생성
- MetaStore: Table, DB, Partitions 의 정보 저장
- Execution Engine: 컴파일러에 의해 생성된 실행 계획을 실행

Hive의 가장 큰 특징은 MetaStore를 관리한다는 것이다. 사용자가 HDFS에 데이터를 저장하면 Hive를 통해 데이터가 테이블 형식으로 다시 저장되는 것이 아니다! 사용자는 HDFS의 데이터를 Hive를 통해 쿼리하기 위해서는 직접 테이블 생성 및 스키마를 설정해주어야 한다. 그러면 Hive는 해당 스키마 정보를 메타스토어에 저장하는 방식이다.