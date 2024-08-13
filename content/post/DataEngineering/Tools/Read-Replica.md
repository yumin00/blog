---
title: "Read Replica 에 대해 알아보자"
date: 2024-08-13T23:08:33+09:00
draft: false
categories :
- DataEngineering
---

이전에, Service DB를 분석에 사용했을 때 발생하는 문제에 대해 정의해보았다([문서](https://yumin.dev/p/service-db-%EC%97%90-%EB%B6%84%EC%84%9D-%EC%BF%BC%EB%A6%AC%EB%A5%BC-%EC%8B%A4%ED%96%89%ED%95%98%EB%A9%B4-%EC%96%B4%EB%96%A4-%EB%AC%B8%EC%A0%9C%EA%B0%80-%EB%B0%9C%EC%83%9D%ED%95%A0%EA%B9%8C/)).
Service DB와 분석 DB를 분리하는 방안 중, Read Replica 에 대해 학습해보고자 한다.  

## Read Replica
![image](https://github.com/user-attachments/assets/b68d2c8b-0d00-4d32-bcde-6b8a1238899e)
먼저, GCP에서 제공하는 Read Replica는 메인 데이터베이스의 데이터를 복제하여 읽기 전용 복제본을 생성하는 기능이다.
Read Replica를 사용하면 기본 인스턴스의 데이터와 기타 변경사항이 거의 실시간으로 읽기 복제본에 업데이트된다.
이를 통해 읽기 요청을 분산시켜 메인 데이터베이스의 부하를 줄이고, 읽기 성능을 향상시킬 수 있다는 특징이 있다.

## 데이터 분석 아키텍처 비교
Service DB에 직접 분석 쿼리를 사용하는 아키텍처는 아래와 같다.

![image](https://github.com/user-attachments/assets/7ef3c10a-2d34-4bd0-8ea6-3856ddeb1fe8)

해당 아키텍처는 다음과 같은 문제를 야기할 수 있으며, 더 자세한 내용은 [이 문서](https://yumin.dev/p/service-db-%EC%97%90-%EB%B6%84%EC%84%9D-%EC%BF%BC%EB%A6%AC%EB%A5%BC-%EC%8B%A4%ED%96%89%ED%95%98%EB%A9%B4-%EC%96%B4%EB%96%A4-%EB%AC%B8%EC%A0%9C%EA%B0%80-%EB%B0%9C%EC%83%9D%ED%95%A0%EA%B9%8C/)에서 확인할 수 있다.

- 서비스 사용자의 경험이 저하될 수 있다.
- CPU와 I/O 리소스를 더 많이 사용하여 데이터베이스 전체 성능을 저하시킬 수 있다.
- 더 심각한 부하가 발생한다면 데이터베이스가 다운될 수 있다.
- 데이터베이스 부하로 인해 트랜잭션이 실패할 수 있으며, 심각한 경우 데이터 손실이 발생할 수 있습니다.

따라서 실제 서비스에 사용되는 DB와 분석을 위한 DB를 분리하기 위해, Read Replica를 사용하여 읽기 요청에 대해서는 복제본을 사용할 수 있다.

![image](https://github.com/user-attachments/assets/359f2960-ee45-40b1-8ce7-8f1905043377)

Read Replica 적용 시, 위 아키텍처와 같이 조회 쿼리에 대해서는 읽기 복제본(Read Replica DB)을 사용할 수 있다.
읽기 작업은 복제된 데이터베이스로 분산시키기 때문에, 원본 데이터베이스의 부하를 줄여 읽기 지연 시간을 단축하고, 전체 응답 시간을 개선할 수 있다.

## 데이터베이스 부하 테스트
### Read Replica 적용 이전 (원본 DB 부하 측정)
<img width="907" alt="image" src="https://github.com/user-attachments/assets/8079aeda-29c7-4e4d-b5c8-106f15491c96">

### Read Replica 적용 후 (원본 DB 부하 측정)
<img width="911" alt="image" src="https://github.com/user-attachments/assets/131f98a3-6431-4178-8f0c-4b0d9c1a377a">

### 분석
Read Replica를 사용함으로써 대시보드에서 발생하는 조회 쿼리는 모두 읽기 복제본 데이터베이스에 요청되기 때문에 원본 데이터베이스에는 부하 및 쿼리 지연 시간이 발생하지 않고 있음을 확인할 수 있다.

## Read Replica 사용 방법
Read Replica를 사용하기 위해서는 직접 읽기 복제본에 연결하여 사용해야 한다. 읽기 복제본은 원본 인스턴스와 다른 IP를 가지고 있기 때문에 이에 직접 연결을 하면 된다.

## 예상 비용 측정
![image](https://github.com/user-attachments/assets/0ffb1b26-78cf-4023-983d-f5a638117445)
GCP 에서 제공하는 예상 가격을 확인해보면, 하루에 약 $2.25로 한 달에 약 $68 가 예상된다.

![image](https://github.com/user-attachments/assets/718daad8-f722-45c0-97da-93782c895340)
하지만 $68 가격은 네트워크와 같은 부수적인 가격이 포함되지 않았고, Google Cloud 가격 책정 문서를 확인하면 읽기 복제본은 독립 실행형 인스턴스와 동일한 요금이 부과되기 때문에 현재 사용하고 있는 Google SQL 가격만큼 비용이 나올 것이라고 예상할 수 있다.

## 결론
Read Replica는 간편하게 DB를 복제할 수 있고, Read Command에 대해서는 원본 DB에 부하를 주지 않을 수 있다는 장점이 있다. 하지만 비용이 인스턴스 하나만큼 더 나오는 것이기 떄문에 비용 측면에서는 단점을 갖고 있다.

따라서, 비용이 적은 방법을 위해서는 Read Replica보다는 원본 DB에 데이터 생성/수정/삭제 를 트리거하여 Bigquery에 적재하는 방식이 더 효율적일 수 있다고 생각한다.




