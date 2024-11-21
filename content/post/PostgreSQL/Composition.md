---
title: "PostgreSQL의 구성에 대해 알아보자"
date: 2024-11-19T19:07:34+09:00
draft: true
categories :
- PostgreSQL
---

오늘부터 postgresql_internals-14를 읽으며 공부한 내용을 작성해보고자 한다.

먼저 PostgreSQL의 구성에 대해 알아보자.

## 데이터 구성
### 데이터베이스
![image](https://github.com/user-attachments/assets/f840261a-42a3-4d93-9f73-e9a53ef19459)

PostgreSQL은 데이터베이스를 관리하는 프로그램이다. PotgreSQL을 하나 생성하면 이를 PostgreSQL Instance 혹은 PosgreSQL Server라고 한다.

그리고 각 데이터베이스를 관리하는 것을 database cluster(PGDATA)라고 한다.

클러스터를 처음 생성하면 세 개의 데이터베이스가 기본으로 생성되어있다.

- template0: 논리적 백업에서 데이터를 복구할 때 사용하는 템플릿 
- template1: 사용자가 해당 클러스터에서 새로운 데이터베이스를 생성할 때 사용하는 템플릿
- postgres: 일반적인 데이터베이스

template0과 template1은 절대로 삭제하면 안된다. 만약 template1을 삭제하면 새로운 데이터베이스를 생성할 수 없다.

Datagrip에서 실제로 template0과 template1이 생성되어 있는지 확인해 보려고 했는데 아무리 찾아도 나오지가 않았다.

![image](https://github.com/user-attachments/assets/796adbee-5522-4962-85d9-648cefa69f28)

클러스터에의 Schema에서 Show template databases를 설정하면 템플릿을 확인할 수 있다.

나의 예상이기는 하지만, template0. template1은 삭제되거나 수정되면 안 되기 때문에 숨겨놓은 것이 아닐까?!
(찾아보니 Datagrip에서는 사용 빈도가 낮은 시스템 데이터베이스는 기본적으로 표시하지 않도록 설정되어 있다고 한다.)

### 시스템 카탈로그
![image](https://github.com/user-attachments/assets/d7a0e1b9-1a84-45aa-9c8c-1aadbfea3b2d)
각 데이터베이스에 있는 테이블, 인덱스와 같은 객체들의 메타데이터는 모두 system catalog라는 테이블에 저장된다.

### 스키마
![image](https://github.com/user-attachments/assets/d7a0e1b9-1a84-45aa-9c8c-1aadbfea3b2d)

스키마는 데이터베이스의 객체를 모두 저장하는 공간이다. 위 사진에서는 test와 같이 새로운 데이터베이스를 생성하면 자동으로 public, information_schema, pg_catalog, pg_toast, pg_temp 스키마가 존재하는 것을 확인할 수 있다.

- public: 가장 기본 스키마로, 데이터베이스를 생성하면 자동으로 public이라는 스키마가 생성된다.
- information_schema: SQL 표준에 따라 제공되는 데이터베이스 메타데이터를 저장하는 스키마이다.
- pg_toast: TOAST 기능과 관련된 객체를 위해 사용된다.
- pg_temp: 임시 테이블을 저장하는 스키마

어떤 객체에 접근하려고 할 때 스키마가 정해져있지 않으면 PosgreSQL은 search_path 매개변수 값에 기반하여 객체를 검색한다.

### search_path
search_path는 검색할 스키마의 순설르 정의하는 매개변수이다. search_path에 설정된 순서대로 스키마를 검색하고, 해당 스키마에서 가장 먼저 일치하는 객체를 선택한다.

```sql
SHOW search_path;
```
이 명령어를 통해 search_path 값을 찾을 수 있는데 만약 설정하지 않았다면 기본 값인 `"$user", public`이 나올 것이다.

```sql
SET search_path TO schema1, schema2, public;
```

그리고 위 명령어를 통해 스키마 검색 순서를 설정할 수 있다.

### 테이블 스페이스
데이터베이스와 스키마는 사용자 입장에서의 데이터 구조이다. Datagrip에서 확인할 수 있듯이 데이터베이스 안에는 스키마가 있고, 스키마 안에는 여러 테이블이 존재한다. 이러한 구조는 논리적인 구조로 사용자 입장에서 데이터가 구분하고 관리하기 쉽다.
이러한 논리적 구조는 실제로 데이터가 디스크에 어떻게 저장되는지는 설명하지 않는다.

테이블 스페이스가 바로 데이터가 물리적으로 어떻게 배치하고 제어할지를 정의한다. 테이블 스페이스를 활용하면 데이터를 분산시키는 것이 가능하고, 자주 사용되는 데이터는 빠른 디스크 저장하고 비교적 접근이 적은 데이터는 느린 디스크에 저장할 수 있다. (이렇게 하면 데이터에 접근하는 속도를 효율적으로 높일 수 있지 않을까..?)

테이블 스페이스 자체는 논리적인 개념이기 때문에 직접 그 형체를 확인할 수 없다. 하지만 기본적으로 클러스터를 생성하면 `pg_global`과 `pg_default`가 생성된다.

- pg_default: 기본 테이블스페이스로, 별도 지정하지 않는 한 모든 테이블과 인덱스가 여기에 저장된다.
- pg_global: PostgreSQL 인스턴스의 전역 시스템 카탈로그를 저장한다.

```sql
SELECT spcname AS tablespace_name,
       pg_tablespace_location(oid) AS location
FROM pg_tablespace;
```

```
tablespace_name
- - - - - - - -
pg_default
pg_global
```

위 SQL문을 통해 클러스터 내의 테이블 스페이스를 조회할 수 있다. 이 경우에는 사용자가 설정한 테이블 스페이스는 존재하지 않고 기본 테이블 스페이스만 존재하는 것을 알 수 있다.

### 릴레이션
PostgreSQL에서 릴레이션이란 관계를 의미한다기 보다는 행과 열로 구성된 객체를 의미한다.

- table: 데이터를 행과 열의 형태로 저장
- index: 