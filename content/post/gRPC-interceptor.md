---
title: "gRPC interceptor"
date: 2023-08-24T17:29:07+09:00
draft: true
---

# gRPC
## gRPC?
gRPC는 구글에서 개발한 RPC 시스템이다. 전송을 위해 TCP/IP 와 HTTP 2.0을 사용하고, IDL(Interface Definition Language)로 protocol buffer를 사용한다.
gRPC에 대해 자세한 내용은 gRPC 포스팅에서 다루고자 한다.

## gPRC 흐름
1. Init
- 시스템이 작동될 때, gRPC 서비스 및 클라이언트를 초기화
- 필요한 구성(ex. 주소, 포트, 인증 정보 등)을 로드하고 gRPC 서버 및 클라이언트 생성

2. gRPC Server
- gRPC 서버는 클라이언트의 RPC 요청을 수신하고 처리
- Protocol Buffers로 정의된 서비스 인터페이스를 구현하여 요청을 처리

3. GateWay
- 클라이언트가 RESTful HTTP 요청을 보내면, Gateway는 이를 gRPC 요청으로 변환하고 해당 서비스에 전달

4. Interceptor & Middleware
- 인터셉터는 gRPC 요청/응답을 가로채는 역할
- Middleware는 일반적으로 여러 인터셉터를 체인으로 연결하는 로직을 의미합니다.
- 로깅, 인증, 트레이싱 등의 공통 로직을 중앙에서 처리할 수 있게 합니다.

5. UnaryServerInterceptor
- Unary RPC에 대한 인터셉터
- 요청을 받아 처리한 후 응답을 반환하기 전에 로직을 실행

gRPC 서비스를 호출할 때, 요청은 여러 인터셉트와 미들웨어를 통과하게 된다. 이 과정에서 로깅, 인증, 권한 확인 등의 로직이 실행된다.
서버는 이러한 로직을 거친 후 요청을 처리하고 응답을 클라이언트에게 반환한다.

gRPC 서비스를 호출하면, Init, Gateway, Interceptor는 콜스택에 쌓이고,
method가 에러를 반환하면, 이는 interceptor로 내려가 redis에 에러가 쌓이는 것이다.