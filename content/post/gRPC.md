---
title: "GRPC"
date: 2023-08-31T17:15:43+09:00
draft: true
---

# gRPC
## IDL(Interface Definition Language)
서버와 클라이언트가 서로 정보를 주고받는 규칙이 프로토콜이라면, IDL은 정보를 저장하는 규칙이다. IDL의 종류에는 3가지가 있다.

1. XML
2. JSON
3. Protocol Buffer

여기서 gPRC에서 사용하는 IDL이 바로 Protocol Buffer이다.


## gRPC란?
gRPC는 구글에서 만든 RPC로, protocol buffer와 RPC를 사용한다.

SSL/TLS를 사용하여 서버를 인증하고 클라이언트와 서버간에 교환되는 모든 데이터를 암호화한다. HTTP 2.0을 사용하여 성능이 뛰어나고 확장 가능한 API를 지원한다.

gRPC에서 클라이언트 응용 프로그램을 서버에서 함수를 바로 호출 할 수 있어 분산 MSA(Micro Service Architecture)를 쉽게 구현 할 수 있다. 서버 측에서는 서버 인터페이스를 
구현하고 gRPC 서버를 실행하여 클라이언트 호출을 처리한다.

## protocol buffer
### protocol buffer 개념
protocol buffer란 IDL(Interface Definition Language)의 종류이다.

> IDL?
> 
> 서버와 클라이언트가 서로 정보를 주고받는 규칙이 프로토콜이라면, IDL은 정보를 저장하는 규칙으로 그 종류에는 XML, JSON, Protocol Buffer가 있다.

프로토콜 버터는 structured data(구조화 데이터)를 직렬화하기 위해 google의 언어로 중립적이기 때문에 확장 가능한 매커니즘이다.

IDL로써 data structure를 정의한 다음, .proto 파일을 protocol buffer compiler(protoc)를 이용해 compile 한다. Compile된 소스 코드를 사용하여 다양한 
데이터 스트림에서 다양한 언어로 다양한 구조의 데이터를 쉽게 읽고 쓸 수 있다.

프로토콜 버퍼는 현재 Java, Python, Objective-C 및 C ++에서 생성 된 코드를 지원합니다. 새로운 proto3 버전을 사용하면 proto2에 비해 더 많은 언어(Dart, Go, Ruby 
및 C #)을 사용할 수 있다.

### protocol buffer 장단점
장점
- 통신이 빠름 : 데이터의 크기가 작기 때문에 더 많은 데이터를 보낼 수 있음
- 파싱을 할 필요가 없음 : JSON 포맷으로 온 데이터는 다시 객체로 파싱해서 사용해야하지만, protocl buffer는 byte stream을 proto file로 읽기 때문에 파싱할 필요가 없음

단점
- 인간이 일긱 불편함 : JSON 포맷의 경우, 사람이 읽기가 편하지만, protocol buffer가 쓴 데이터는 proto 파일이 없으면 무슨 의미인지 알 수 없음.
  - 때문에 모든 클라이언트는 proto 파일을 가지고 있어야 해서 외부 API로 쓰이기에도 문제가 있다.
  - 그래서 내부 서비스 간의 데이터 교환에 주로 사용된다.
  - API Gateway를 이용하여 REST API로 바꿔주는 방식도 존재한다.