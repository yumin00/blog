---
title: "JWT에 대해 알아보자"
date: 2023-11-07T23:45:39+09:00
draft: false
categories :
- Network
---

# JWT
## 개념
JWT(Json Web Token)은 인터넷 표준 인증 방식이다. 인증에 필요한 정보를 Token에 담아 암호화시켜 사용하는 토큰이다.

JWT의 중요한 점은 서명된 토큰이라는 것이다. 공개/개인 키를 쌍으로 사용하여 토큰에 서명할 경우 서명된 토큰은 개인키를 보유한 서버가 이 서명된 토큰이 정상적인 토큰인지 인증할 수 있다는 것이다.

## 구조`
- Header
- Payload
- Signature`

### Header
헤더에는 토큰의 타입이나 서명 생성에 어떤 알고리즘이 사용되었는지 저장한다.

```json
{
  "type": "JWT",
  "alg": "HS512"
}
```

### Payload
페이로드에는 보통 Claim이라는 사용자에 대한, 혹은 토큰에 대한 property를 key-value 형태로 저장한다.

- iss: 토큰 발급자
- sub: 토큰 제목 - 토큰에서 사용자에 대한 식별 값이 됨
- aud: 토큰 대상자
- exp: 토큰 만료 시간
- nbf: 토큰 활성 시간
- iat: 토큰 발급 시간
- jti: JWT 토큰 식별자 - issuer가 여러 명일 때 이를 구분하기 위한 값

페이로드에는 민감한 정보를 담지 않아야 한다. 헤더와 페이로드는 json이 디코딩되어 있을 뿐, 특별한 암호화가 되어 있지 않기 때문에 누구나 jwt를 가지고 디코딩할 수 있기 때문이다.

### Signature
시그니처는 서버에 있는 개인키로만 암호화를 풀 수 있으며, 다른 클라이언트는 임의로 복호화할 수 없다.

### JWT
JWT를 만드는 방법은 다음과 같다.

- base64UrlEncode(Hedaer): json 형식의 Header를 Base64Url로 인코딩
- base64UrlEncode(Payload): json 형식의 Payload 를 Base64Url로 인코딩
- signature : <base64UrlEncode(header) + "." + base64UrlEncode(payload), secretKey> 를 header에서 지정한 알고리즘으로 생성

JWT = base64UrlEncode(Hedaer) + "." + base64UrlEncode(Payload) + "." + signature

## JWT는 왜 사용하는걸까?
### stateful 인증
- 서버가 사용자의 로그인 상태를 가지고 있어야 한다. 즉, 사용자의 로그인 여부를 DB에 저장해야 한다.

### stateless 인증
- 서버가 사용자의 로그인 상태를 가지고 있지 않는다.
- 대신, 각 요청마다 인증 정보를 받아 인증이 된 사용자인지 판단한다.

stateful 인증을 사용하면, 서버는 사용자의 로그인 상태를 DB에 저장해야 한다. 각 요청마다 사용자의 로그인 상태를 DB를 통해 가져와서 인증 여부를 확인해야 한다. 그렇게 되면 성능에 영향을 줄 수 있다.

따라서 stateless 인증을 사용하면 DB를 거칠 필요없이, 성능에 영향을 주지 않고 더 빠르게 요청/응답이 가능해진다. 그래서 JWT가 사용되는 것이다.

하지만, 세션과 달리 클라이언트에 유저의 정보가 포함된 토큰이 저장되는 것이기 때문에 민감한 정보를 담아서는 안된다.  

현재 진행중인 프로젝트에서는 인증 방식을 다음과 같이 다루고 있다.
1. 사용자가 로그인하면 서버는 refresh token과 access token을 생성한다.
- refresh token : 랜덤 문자열로 DB에 저장
- access toekn : JWT

2. API를 호출할 때마다 클라이언트는 header에 access token 을 포함하여 요청하고, 서버는 이를 검증한 뒤 응답을 보낸다.
3. access token이 만료되면 refresh token을 검증한 뒤, 새로운 access token을 발급한다.


다음에는 이러한 JWT 인증을 go 에서 어떻게 구현할지에 대해 이야기해보고자 한다.