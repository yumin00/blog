---
title: "Postman에서 Mock Server 만들기"
date: 2024-01-03T23:50:20+09:00
draft: true
categories :
- Documentation
---

# Postman 에서 Mock Server 만들기
사이드 프로젝트를 진행하면서, 클라이언트 개발자에게 API 문서와 API Mock 서버를 전달해야하는 일이 발생했다. 클라이언트 개발자가 API 개발이 모두 완료될 때까지
기다릴 수 없기 때문에, API Mock 서버를 제공함으로써 클라이언트 개발자는 좀 더 빠르고 수월하게 개발할 수 있다.


Postman으로 Mock Server를 만들어 클라이언트 개발자에게 전달하는 과정은 처음으로 진행해보기 때문에 이에 대해 공부해보고자 한다.

## Mock Server란 무엇일까?
Mock 서버란 가짜 서버라고 이해할 수 있다. 특정 Request를 날리면 가짜 데이터를 보내주는 가상 서버를 Mock 서버라고 한다.

클라이언트 개발자는 API가 모두 완성될 때까지 기다리지 않고, Mock 서버를 통해 쉽게 테스트를 진행해볼 수 있다.

## Postman을 통해 Mock Server 만들기
이제 직접 포스트맨을 사용하여 Mock 서버를 만들어보고자 한다.

1. My Workspace에서 `New` 버튼을 클릭하여 Collection을 만들어 준다.