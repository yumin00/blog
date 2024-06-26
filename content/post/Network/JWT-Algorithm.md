---
title: "JWT Signature 알고리즘에 대해 알아보자"
date: 2023-11-14T20:33:19+09:00
draft: false
categories :
- Network
---

# JWT Signature 알고리즘
## JWT Signature 개념
JWT 시그니쳐는 서명에 대한 정보로, 헤더와 페이로드를 인코딩한 값을 합치고, 특정 알고리즘과 특정 키로 암호화되어 있는 값이다. 암호화되어 있기 때문에 복호화를 하지 않으면 볼 수 없다.

헤더와 페이로드는 암호화된 값이 아니라, 인코딩되어 있다.
**헤더와 페이로드는 보안이 취약하기 때문에 시그니처를 추가하여 위변조를 판별할 수 있다.(너가 만든게 맞냐고 체크)**

## JWT 암호화 방식
### 1. 대칭키 암호화
암호화, 복호화 키가 같을 경우 대칭키 암호화 방식이라고 한다.

- 같은 키를 사용하여 암호화, 복호화하기 때문에 속도가 빠르다.
- 대표적으로 HMAC 암호화 알고리즘이 있다.
- 값에 SHA256를 적용해서 해싱 후 private key(대칭키 역할)로 암호화한다.
- private key를 알고 있는 서버만 JWT를 복호화할 수 있다.
- auth 서버가 없으면 대칭키가 편할 수 있음.

### 2. 비대칭키 암호화
암호화, 복호화 키가 다른 경우 비대칭키 암호화 방식이라고 한다.

- 속도가 느리지만, 대칭키 암호화에 비해 안전하다.
- 대표적으로 RSA 암호화 알고리즘이 있다.
- 값에 SHA256을 적용해서 해싱 후 비밀키로 암호화한다.
- 공개키는 공개적으로 제공한다. 어떤 서버든 공개키를 통해 JWT를 복호화할 수 있다.

## SHA256 알고리즘
HS256, RS256 알고리즘에서 공통적으로 사용되는 SHA256 알고리즘에 대해 알아보자. 

- SHA(Secure Hash Algorithm)의 한 종류로 256비트로 구성되며 64자리 문자열을 반환
- 블록체인에서 가장 많이 사용
- 2^256 만큼 경우의 수를 만들 수 있음
- 많은 시간이 소요될 정도로 큰 숫자이기 때문에 충돌로부터 비교적 안전하다고 평가됨
- 해쉬 알고리즘은 복호화가 불가능하다. 즉, 해쉬 알고리즘 종류 중 하나인 SHA256 알고리즘은 복호화가 불가능하다.

## HS256 알고리즘 [대칭키 암호화]
- JWT 암호화 알고리즘으로 많이 사용한다.
- HS256 알고리즘 = HMAC + SHA256
- 메시지(ex.JWT의 페이로드)와 비밀키를 함께 사용하여 HMAC 생성
- SHA-256 해시 함수를 사용하여 메시지와 비밀키의 조합으로부터 고유한 해시 값을 생성
- 공유된 비밀키가 안전하게 관리될 수 있는 환경에서 효과적
- 비밀키가 노출될 경우, 보안에 위협받을 수 있기 때문에 비밀키 관리가 중요하다.

### HS256 동작 방식
- JWT 생성: 헤더는 알고리즘과 토큰 타입을 명시하고, 페이로드는 클레임(예: 사용자 식별 정보)을 포함
- 비밀키 설정: 서버는 비밀키를 생성하고 안전하게 보관해야 한다. 이 비밀키는 JWT의 서명 생성과 검증에 사용된다.
- 시그니처 생성: HS256 알고리즘을 사용하여 JWT의 헤더와 페이로드를 결합한 문자열에 시그니처를 생성
- JWT 전송 및 검증: 생성된 JWT는 클라이언트에게 전송한다. 클라이언트는 요청 시 JWT를 포함하여 서버에 전송한다. 서버는 비밀키를 사용하여 시그니처를 검증하고, 검증이 성공하면 요청을 처리한다.

## RS256 알고리즘 [비대칭키 암호화]
- RS256은 공개키와 비밀키를 사용한다.
  - 서명 생성 - 개인키
  - 서명 검증 - 공개키
  - 수신자가 공개키로 시그니처를 검증했다는 것은, 개인키를 가진 발신자가 시그니처를 생성했고 메시지가 변경되지 않았음을 의미한다. 
- RSA 암호화 방식으로, 복잡한 수학적 연산을 기반으로 하기 때문에 이를 통해 생성된 시그니처는 메시지의 무결성과 인증을 보장한다.

### RS256 동작 방식
- 키 생성: 공개키와 비밀키 생성. 비밀키는 서버에 안전하게 보관되며, 공개키는 수신자에게 공유된다.
- JWT 생성: 헤더에는 사용된 알고리즘(RS256)이 명시되고, 페이로드에는 클레임이 포함된다.
- 시그니처 생성: 비밀키를 사용하여 JWT의 헤더와 페이로드에 서명한다. 이 시그니처는 JWT의 마지막 부분을 형성한다.
- JWT 전송 및 검증: 생성된 JWT는 클라이언트에게 전송된다. 클라이언트는 요청 시 JWT를 포함하여 서버에 전송한다. 수신자는 공개키를 사용하여 시그니처를 검증한다.