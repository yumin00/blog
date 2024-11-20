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
