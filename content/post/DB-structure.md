---
title: "DB Structure"
date: 2023-11-18T23:53:18+09:00
draft: true
categories :
- Network
---

# Auth를 위한 DB 구조
## 1. 인증 / 인가 프로세스
팀프로젝트에서, Auth를 위한 DB 구조를  정해야할 필요가 있다.

먼저, Auth에는 authZ와 authE가 있다.

> - authZ : 인증 - 서버에 접근하는 인증
> - authE : 인가 - 회원과 비회원을 체크하기 위한 인가

## authZ
서버에 접근하는 인증에서는 stateless한 앱의 특성상 토큰 방식이 어울릴 것이라고 생각하여 그 중 가장 친숙한 JWT를 선택했다.

JWT 암호화 알고리즘에는 대칭과 비대칭 방식으로 나눌 수 있는데, 대칭키 방식의 HS256 알고리즘을 사용하기로 했다.

## authE
인가에서는 회원임을 판단하는 기준을 설정해야 한다. 인증에서 signature를 사용했다면 인가에서는 payload를 사용한다.

payload에서 aud를 통해 회원/비회원 여부를 판단하고자 한다.

[authZ / authE flow]

인증/인가 flow는 다음과 같다.

<img width="1277" alt="image" src="https://github.com/yumin00/blog/assets/130362583/2e522418-eeac-45ed-a471-980114a5f81c">

- 인증
  - 인증을 위해서 게스트토큰을 사용한다. (게스트토큰은 갱신 X)
  - 비회원은 게스트 토큰을 사용하고, userid를 갖는다.
  - 게스트토큰으로 API를 요청하고, 서버는 게스트토큰의 유효성을 검증한 뒤 유효할 경우 response를 준다.
- 인가
  - 회원가입 후, 로그인을 했을 경우 해당 user에 대해서 리프레시 토큰을 발급한다. 이때, 리프레시 토큰은 DB에 저장한다.
  - 로그인 이후 API를 요청할 때, 해당 리프레시 토큰을 검증 후 액세스 토큰을 발급한다.
  - 로그인을 할 때마다 리프레시 토큰을 재발급한다.
  - 클라이언트는 리프레시 토큰/액세스 토큰을 사용하여 API를 요청하고, 서버는 토큰의 유효성을 검증한 뒤 유효할 경우 response를 준다.
  - 액세스 토큰이 만료됐을 경우(Unauthorized), 리프레시 토큰을 유효성을 확인하여 액세스 토큰을 재발급해준다.

해당 플로우를 위해 auth를 위한 user db는 다음과 같이 구성해보았다.

<img width="500" alt="image" src="https://github.com/yumin00/blog/assets/130362583/2fbd530a-a862-4656-8e70-5d44d7979b75">

### 액세스 토큰
- stateless

### 리프레시 토큰
- stateful
- 서버에서 user 별로 random string으로 저장
- 갱신 요청 시, 기존 리프레시 토큰과 비교 후 새로운 리프레시 토큰 발급
- 로그아웃했을 때, 새로 발급이 필요함
- 만료기간이 없음

## 2. JWT 토큰 구조
- Header 세부 정보
```json
{
  "type": "JWT",
  "alg": "HS512"
}
```
- Payload 세부 정보
```json
{
  "iat": 발급한 시간,
  "iss": 토큰 발급자,
  "exp": 만료기간(5분),
  "aud": user_type(회원/비회원),
  "uid": user_id,
}
```
- Signature 생성 방법
  - HS256 알고리즘 사용
  - 대칭키(문자열)를 base64 encode 2번

## 3. 토큰 발급 및 검증
- 대칭키를 사용한 토큰 발급 과정
- 토큰 검증 과정
- 토큰 만료 및 갱신 과정

## 4. 보안 고려사항
- 대칭키 관리 방법
- 토큰 보안 취약점 및 대응 방안
  - 액세스토큰가 탈취되면, 만료기간까지 어떻게 할 방법이 없음 -> 액세스토큰의 만료 기간을 짧게 가져가야한다. 따라서, 액세스 토큰의 만료 기간을 5분으로 설정한다.

## 5. 실제 구현 예시

## 6. 향후 개선 방향