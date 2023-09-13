
---
title: "GDG 송도"
date: 2023-08-26T14:04:28+09:00
draft: true
---

# Feature Toggles == Feature Flag
코드를 수정하지 않고, 시스템 동작을 변경하는 기술.

## 카테고리
### 1. 릴리즈 플래그
- 불안전하고 테스트 되지 않은 코드를 숨긴채로 Production에 배포 가ㄴ능
- 완료되지 않은 기능을 숨기기
- main branch에 바로 merge
- TBD 전략을 사용하는 사례가 많아지면서, 이를 위해 사용할 수 있는 플래그
- 기능이 완성되면 사라지기 때문에 생명주기가 짧음

### 2. 실험 플래그
- A/B 테스트, 개인화 등을 위해 사용되는 플래그
- 광고배너의 상/하단 노출 여부, 하단 탭배치 등 데이터 기반 의사결정을 하기 위한 토글

### 3. 운영 플래그
- 외부 시스템 이중화 운영에 사용 (지도 검색 API, 결제업체, 본인인증 업체 등)
- 혹은 성능상 이슈가 될 피쳐를 상황에 따라 비활성화
- 특정 이슈로 사용자가 급격히 몰릴 것이 예상될 때, 이슈가 생길 수 있는 추천 패널이나 개인화된 뷰 기능 OFF
- 일종의 manual한 circuit breaker
- 생명 주기가 길다

### 4. 권한 플래그
- 임의로 지정한 특정 사용자에게만 feature 노출
- ex) 카카오톡 실험실 기능, 임직원 테스트
- 카나리는 무작위, 권한은 지정

## 구현
- GrowthBook

### 토글 포인트 / 토글 라우터 / 토글 설정
- 토글 포인트 : 어떤 토글이 true일 때, 여기를 타라는 if문 / 구분하는 코드
- 토글 라우터 : flag 값이 켜져있는지 아닌지 판단하는 로직 / 구분하기 위한 객체나 수단
  - id 특정 케이스인 경우와 같은 로직
- 토글 설정


### 구현 원칙
- on : 신규 feature
- 토글 포인트 / 라우터 분리
  - 토글 포인트는 왜 토글이 켜졌는지 꺼졌는지는 알 필요가 없음
  - if - else 를 통해 어떤 메소드를 실행할지는 토글 포인트의 역할
- CNCF 샌드박스 프로젝트

## 단점
- 관리해야 할 코드가 2배로 늘어남.
- 릴리즈 플래그는 관리가 필수


## RDB vs NoSQL
### 1. RDB
- 관계형 데이터베이스
- 2차원 테이브 형태
- 트렌젝션과 ACID(원자성, 일관성, 고립성, 지속성) 지원

### 2. NoSQL
- key-value, document, colum family db, graph db 등의 비정형화된 데이터 구조를 갖는 db
- 데이터 간의 관계를 정의하지 않음

## in-memory
- 주메모리(ram)에 데이터를 적재하여 사용
- 위로 갈수록 성능이 좋고, 용량은 작고, 비용이 비쌈
- cpu에서 ram에 접근할 수 있으므로 하드디스크에 비해 메모리 접근 속도가 1,000배 정도 빠름
- 휘발성이므로 전원공급이 안될 경우 데이터가 유실될 수 있음
- ex ) redis, memcached

## Redis vs Memcached
### Redis
- single thread 사용으로 한번에 1개의 명령어만 처리 가능 - O(N)
- 많은 자료구조 제공
- 클러스터 모드 제공
- 메모리 관리 필수 


# 모듈러 모놀리스
## 모놀리스 vs 마이크로 서비스 [개념]
### 1. 모놀리스
- 하나의 어플리케이션 안에 다양한 기능을 하는 요소들이 다같이 모여있음
- 모든 기능이 하나의 어플리케이션에 집중
- 고전적인 방식
- 제한된 확장성

### 2. 마이크로서비스 아키텍처
- 각각 기능을 나눠서 처리
- db도 여러개로 구성
- 기능들을 여러 개의 어플리케이션으로 흩어짐
- 큰 규모의 기업들이 선택하는 구조
- 이론상 무한한 확장성

## 모놀리스 va 마이크로서비스 [활용]
### 1. 모놀리스
- 하나의 서버 배포
- method call 통신
- 하나의 데이터베이스
- 트렌젝션
- 작은 규모의 조직에 적합
- 낮은 운영 비용
- 일부 구성 교체 불가

### 2. 마이크로 서비스
- 여러 서버 배포
- network call 통신
- 여러개의 데이터베이스
- 최종적 일관성
- 큰 규모의 조직에 적합
- 높은 운영 비용

## 어떤 아키텍처를 선택해야 할까?
모듈화 되지 않은 모놀리스
- 비즈니스별 물리적/개념적 분리 x

모듈러 모놀리스
- 비즈니스별 물리적 분리 X
- 비즈니스별 개념적 분리 O

마이크로서비스
- 비즈니스별 물리적/개념적 분리 O

## 모듈러 모놀리스
- 바운디드 컨텍스트(Bounded context)는 모듈을 나눌 때도 유용
- 순환 참조를 피하자