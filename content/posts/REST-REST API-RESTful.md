---
title: "REST vs REST API vs RESTful"
date: 2023-08-23T17:12:43+09:00
draft: true
---

# REST vs REST API vs RESTful
## 1. REST
### REST 정의
**REST(REpresentational State Transfer)** 이란 자원의 이름으로 구분하여 해당 자원의 상태를 주고받는 모든 것을 의미한다. 즉, 자원의 표현에 의한 상태 전달이다.
- 자원의 표현
  - 자원 : 해당 소프트웨어가 관리하는 것
  - 자원의 표현 : 그 자원을 표현하기 위한 이름
  - ex) DB의 '학생 정보'가 **자원** 일 때, 'students'가 **자원의 표현**
- 상태 전달
  - 데이터가 요청되어지는 시점에 자원의 상태(정보)를 전달
  - JSON / XML 을 통해 전달

### REST 구성요소
1. 자원 : HTTP URI
2. 자원에 대한 행위 : HTTP METHOD
3. 자원에 대한 행위의 내 : HTTP Message Pay Load

=> HTTP URI를 통해 자원(Resource)을 명시하고, HTTP METHOD(POST, GET, PUT, DELETE)를 통해 해당 자원에 대한 CRUD Operation을 적용하는 것을 의미한다.
즉, 자원 기반 구조의 설계 중심으로 HTTP METHOD를 통해 자원을 처리하도록 설계된 아키텍처를 의미한다.

- CRUD Operation
  - Create
  - Read
  - Update
  - Delete
  - HEAD : header 정보 조회

### REST 특징
- 네트워크 상에서 Client와 Server 사이의 통신 방식 중 하나
- HTTP URI 를 통해 자원(Resource)을 명시하고, HTTP Method(POST, GET, PUT DELETE) 를 통해 해당 자원에 대한 CRUD Operation을 적용하는 것을 의미

### REST 장단점
- 장점
  - HTTP 프로토콜의 인프라를 그대로 사용하기 때문에 REST API 사용을 위한 별도의 인프라를 구축할 필요가 없다.
  - REST API 메시지가 의도하는 바를 명확하게 나타내므로 의도하는 바를 쉽게 파악할 수 있다.
- 단점
  - 표준이 존재하지 않음
  - 사용할 있는 METHOD가 CRUD로 한정

  
## 2. REST API?
### REST API  개념
REST API는 REST의 원리를 따르는 API를 말한다. 

### REST API 설계 규칙
- URI는 자원의 정보를 표시해야 한다.
  - resource는 동사보다 명사, 대문자보다 소문자를 사용
  - resource의 도큐먼트 이름은 단수 명사 사용
  - resource의 컬렉션 이름은 복수 명사 사용
  - resource의 스토어 이름은 복수 명사 사용
- 자원에 대한 행위는 HTTP Method로 사용한다.
  - URI에 HTTP Method가 들어가면 안됨
- 언더바 대신 하이픈을 사용한다.
- 파일확장자는 URI에 포함하지 않는다.

## 3. RESTful
### RESTful 개념
REST API를 제공하는 웹 서비스를 RESTful 하다고 할 수 있다. RESTful API는 REST API 설계 가이드를 따라 API를 만드는것 입니다.


### RESTful 목적
이해하기 쉽고 사용하기 쉽게 REST API를 만드는 것이다. API를 RESTful 하게 만들어서 API의 목적이 무엇인지 명확하게 하기 위해 RESTful 함을 지향한다.