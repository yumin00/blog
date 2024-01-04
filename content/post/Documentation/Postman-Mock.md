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
<img width="577" alt="image" src="https://github.com/yumin00/blog/assets/130362583/c7a4690e-0c2e-4670-be72-9986c4c5c053">

2. 해당 Collection에서 우클릭하여 `Mock collection` 을 생성한다.
<img width="297" alt="image" src="https://github.com/yumin00/blog/assets/130362583/96b2f764-13e8-4ef4-aef1-4092cf1180bb">

3. Mock 서버 정보를 입력한 후 Mock 서버를 생성한다.
<img width="892" alt="image" src="https://github.com/yumin00/blog/assets/130362583/58041678-f865-4030-9677-7c943c208b35">
<img width="818" alt="image" src="https://github.com/yumin00/blog/assets/130362583/db62f547-0161-4e76-85cc-14e8e0cea4c0">

해당 과정을 통해 Mock 서버를 생성할 수 있다. 그 다음으로는 클라이언트 개발자에게 전달하고싶은 API를 Mock 서버에 등록해주어야 한다.

1. 우클릭을 통해 `Add Request`로 새로운 API를 생성한다.
<img width="295" alt="image" src="https://github.com/yumin00/blog/assets/130362583/0df624af-fc29-4cbc-8fed-84fe4b361b1c">

2. `Add example` 을 통해 생성하고자 하는 API에 대한 example을 작성 후 저장한다. 이때 url은 Mock 서버 생성 후 제공받은 url을 사용하여 작성해야 한다. response 
   값에 대한 
   예시도 함께 작성해준다. 
<img width="274" alt="image" src="https://github.com/yumin00/blog/assets/130362583/e7f71787-e9fb-4500-ae9f-2336ebaa1dc2">

3. Mock 서버에 요청을 날릴 경우, 다음과 같이 설정한 예시 값으로 나오는 것을 확인할 수 있다.
<img width="890" alt="image" src="https://github.com/yumin00/blog/assets/130362583/01996c98-24f8-4932-91d6-8f9e75c50b41">

이렇게 Mock 서버를 생성함으로써, 백엔드 서버에서 API 요청 시 예상되는 요청/응답 값을 모두 설정해보았다. 이를 통해 백엔드 개발이 완료되지 않아도 클라이언트 개발을 진행할 수 있다.