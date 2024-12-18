---
title: "GORM을 사용하여 PostgreSQL에 연결하기"
date: 2024-12-18T19:41:14+09:00
draft: true
categories :
- PosrgreSQL
---

## GORM을 사용하여 PostgreSQL에 연결
Go를 사용하여 Postgres에 연결할 때 ORM인 GORM을 사용할 수 있다.

Postgres에 연결하는 방법은 간단하다.

```go
package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
    }
	db.Table("users").Select("name, age")
}
```

### 커넥션 풀
GORM은 database/sql을 사용하여 커넥션 풀을 관리한다.
```go
postgreDB := db.DB()
postgreDB.SetMaxIdleConns(10)
postgreDB.SetMaxOpenConns(100)
postgreDB.SetConnMaxLifetime(time.Hour)
```

- SetMaxIdleConns: DB 연결 풀에서 유지할 수 있는 유휴(Idle) 연결의 최대 개수를 설정하는 것이다.
- SetMaxOpenConns: DB와 동시에 열 수 있는 최대 연결 개수를 설정
  - 값이 너무 낮으면 동시 요청이 많을 때 연결을 대기하는 시간이 길어질 수 있다
  - 값이 너무 크면 데이터베이스 서버에 부하가 발생하거나 과도한 연결 관리로 문제가 생길 수 있다.
- SetConnMaxLifetime: 하나의 연결이 재사용되기 전까지의 최대 지속시간

## Cloud Run과 Cloud SQL
Cloud Run에서 Unit Socket으로 VPC를 사용하여 Cloud SQL에 접근할 때, 한 인스턴스당 100개의 연결이 가능하다.
![image](https://github.com/user-attachments/assets/6bc71c20-00e5-45f6-b713-7e9bec9c8a4a)

이때 한 인스턴스에 100개를 초과하는 요청이 들어왔다면 Cloud SQL에 접근하기 위해서는 어떻게 해야할까?