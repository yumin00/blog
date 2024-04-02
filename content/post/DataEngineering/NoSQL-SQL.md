---
title: "NoSQL vs SQL"
date: 2024-04-02T17:43:38+09:00
draft: true
categories :
- DataBase
- DataEngineering
---

# NoSQL vs SQL

# 데이터베이스(Database, DB)
SQL과 NoSQL에 대해 알기 전, 데이터베이스에 대한 이해가 필요하다.
## 데이터베이스란?
데이터베이스란 데이터의 집합으로, 데이터를 쉽게 접근, 업데이트 등을 할 수 있도록 한다.  

## 데이터베이스 관리 시스템(Database Management System, DBMS)
데이터베이스를 관리하기 위해 데이터베이스 관리 시스템을 사용한다. 데이터베이스란 원론적으로는 데이터베이스의 모음을 말하지만, 데이터를 조직하는 유형별로 혹은 DBMS 솔루션까지 데이터베이스라고 칭한다.

----

데이터베이스의 유형에는 관계형 데이터베이스(Relational DataBase Management System, RDBMS)와 NoSQL 데이터베이스가 있다. 관계형 데이터베이스는 SQL을 사용하고 NoSQL 데이터베이스에서는 NoSQL을 사용한다.

지금부터 NoSQL과 SQL의 차이점에 대해 알아보자

## SQL
![image](https://github.com/yumin00/blog/assets/130362583/6fb07414-e667-4673-a461-bd93ad1860c2)

관계형 데이터베이스란, 고정된 열과 행을 구성된 테이블에 데이터를 저장하는 방식을 말한다. 관계형 데이터베이스를 사용할 때는 테이블의 구조와 데이터 타입을 명시하여 사용해야 한다.

관계형 데이터베이스를 관리하는 시스템을 RDBS라고 한다.

> 관계형 데이터베이스
> 
> - MySQL
> - Oracle
> - SQlite
> - MariaDB
> - PostgreSQL

각 시스템의 특징은 다음에 이어서 공부해보고자 한다. :)

SQL은 RDBMS의 데이터를 관리하기 위해 설계된 특수 프로그래밍 언어다. 즉, RDBMS의 전용 프로그래밍 언어라고 할 수 있다!

![image](https://github.com/yumin00/blog/assets/130362583/ec1e52d9-d377-46da-add0-c53605f9c1ec)
관계형 데이터베이스에서는 테이블마다 명확한 스키마가 있고, 데이터 중복을 방지하기 위해 각 테이블은 관계를 가진다.

## NoSQL
![image](https://github.com/yumin00/blog/assets/130362583/2430f20f-ac55-449c-af32-d9dc8637417c)

비관계형 데이터베이스란, 관계형 데이터베이스를 제외한 나머지를 말한다. 따라서 비관계형 데이터베이스를 NoSQL이라고 한다.
관계헝 데이터베이스에서는 데이터를 입력할 때 스키마에 맞춰 입력해야 하지만, 비관계형 데이터베이스에서는 데이터를 읽어올 때만 스키마에 따라 데이터를 읽어온다. 이러한 방식을 `schema or read`라고 한다.

![image](https://github.com/yumin00/blog/assets/130362583/0d39b11d-eea2-421b-aaf6-1baee492f13a)
NoSQL에서는 스키마도 없고, 관계형도 없다. 그러면 조인하고 싶을 때, NoSQL에서는 어떻게 할 수 있을까? 위 그림처럼, 한 컬렉션에서 조인하고자 하는 데이터의 일부분을 복제해서 넣어주어야 한다.
그러면, 컬렉션의 일부분에 속하는 데이터를 산출할 수 있다.


