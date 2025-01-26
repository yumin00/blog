---
title: "사이드 프로젝트에서 사용할 DB 골라보기"
date: 2025-01-26T23:37:36+09:00
draft: true
categories :
- DB
---

이번에 사이드 프로젝트를 진행하면서 필요한 DB를 골라본 과정에 대해 글을 작성해보고자 한다.

먼저 사이드 프로젝트의 데이터 구조는 정규화가 필요하기 때문에 RDBMS를 선택하고자 했다.

## 1. AWS RDS vs AWS Aurora 
프리티어로 AWS를 사용하면 저렴한 운영이 가능하기 때문에 먼저 AWS에서 사용해볼 수 있는 DB에 대해 조사해봤다.

### AWS RDS
- 관리형 관계형 데이터베이스 서비스
- MySQL, PostgreSQL, MariaDB와 같은 DB를 운영해줌
- 자동 백업, 복구, 소프트웨어 패치 같은 운영 작업을 AWS가 관리해줌
- 비교적 저렴함


### AWS Aurora
- AWS가 자체적으로 만든 MySQL, PostgreSQL과 호환되는 고성능 관계형 데이터베이스 서비스
- MySQL은 최대 5배, PostgreSQL은 최대 3배 빠름
- RDS보다는 비싸지만 높은 성능을 제공함.

현재 진행하는 사이드 프로젝트는 아직 운영 단계가 아니고 저렴한 가격의 SaaS를 사용하고 싶기 때문에 AWS에서 운영한다면 Aurora보다는 RDS가 더 좋은 선택일 것 같다.




## 2. Oracle Cloud Autonomous Database
[Oracle Cloud](https://www.oracle.com/kr/cloud/free/)는 무료로 사용할 수 있는 "Always Free" 리소스를 제공하며 여기에는 데이터베이스 서비스도 포함된다고 한다.

- Autonomous Database (ATP 또는 ADW)
  - ATP(Autonomous Transaction Processing): 트랜잭션 중심의 워크로드에 적합. 
  - ADW(Autonomous Data Warehouse): 분석 중심의 워크로드에 적합. 
  - 제공 사항:
    - 2개의 Autonomous Database 인스턴스 
    - 각각 최대 20GB의 스토리지 
    - 자동 백업 및 고가용성 지원 
    - 머신 러닝 기반의 자동 튜닝과 패치 

