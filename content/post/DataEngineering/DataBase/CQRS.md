---
title: "CQRS 패턴에 대해 알아보자"
date: 2024-05-02T20:10:04+09:00
draft: true
categories :
- DataBase
- DataEngineering
---

# CQRS (Command and Query Responsibility Segregation)
CQRS 패턴이란, 데이터 저장소로부터의 명령와 조회 작업을 분리하는 패턴을 말한다.

그렇다면, 명령과 조회는 무엇이고 명령과 조회는 왜 분리해야할까?

## 명령과 조회
데이터베이스에는 생성(Create), 조회(Read), 갱신(Update), 삭제(Delete)할 수 있다. 여기에서, 조회를 제외한 생성, 갱신, 삭제가 명령이다.
