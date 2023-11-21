---
title: "DB Structure"
date: 2023-11-18T23:53:18+09:00
draft: true
categories :
- Network
---

# Auth를 위한 DB 구조
팀프로젝트에서, Auth를 위한 DB 구조를  정해야할 필요가 있다.

먼저, Auth에는 authZ와 authE가 있다.

> - authZ : 인증 - 서버에 접근하는 인증
> - authE : 인가 - 회원과 비회원을 체크하기 위한 인가

## authZ
서버에 접근하는 인증에서는 stateless한 앱의 특성상 토큰 방식이 어울릴 것이라고 생각하여 그 중 가장 친숙한 JWT를 선택했다.

JWT 암호화 알고리즘에는 대칭과 비대칭 방식으로 나눌 수 있는데, 대칭키 방식의 HS256 알고리즘을 사용하기로 했다.

[authZ flow]
인증 flow는 다음과 같다.

## authE
인증에서 JWT의 Signature를 이용했다면, 인가에서는 Payload를 이용하기로 했다.

JWT의 특성을 이용하여 검증이 완료된 토큰의 경우, 믿을 수 있기 때문에 Payload에 있는 userId 혹은 roleId를 통해서 
이 토큰이 접근할 수 있는 있는지를 판단할 수 있도록 인가를 진행했다.

[authE flow]