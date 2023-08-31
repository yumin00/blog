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
gRPC는 구글에서 만든 RPC로, protocol buffer와 RPC를 사용합니다.

SSL/TLS를 사용하여 서버를 인증하고 클라이언트와 서버간에 교환되는 모든 데이터를 암호화합니다. HTTP 2.0을 사용하여 성능이 뛰어나고 확장 가능한 API를 지원합니다.

gRPC에서 클라이언트 응용 프로그램을 서버에서 함수를 바로 호출 할 수 있어 분산 MSA(Micro Service Architecture)를 쉽게 구현 할 수 있습니다. 서버 측에서는 서버 인터페이스를 구현하고 gRPC 서버를 실행하여 클라이언트 호출을 처리합니다

### 1. protocol buffer란?
protocol buffer란 IDL(Interface Definition Language)의 종류이다.

> IDL?
> 
> 서버와 클라이언트가 서로 정보를 주고받는 규칙이 프로토콜이라면, IDL은 정보를 저장하는 규칙으로 그 종류에는 XML, JSON, Protocol Buffer가 있다.

### 2. SSL/TLS 란?
