---
title: "OAuth에 대해 알아보자"
date: 2023-10-04T23:40:43+09:00
draft: true
categories :
- Security
---

# OAuth
## 1. OAuth 등장 배경
우리는 우리의 서비스에서 사용자의 네이버, 카카오와 같은 아이디와 비밀번호를 입력받아 관련된 서비스를 제공할 수 있다. 로그인을 할 때, 사용자로부터 아이디 로그인을 입력 받으면 해당 정보를 서비스에 저장하고 활용한다. 이 방법은 안전할까? 사용자들은 처음 보는 서비스를 신뢰하여 아이디와 비밀번호를 입력할 수 있을까?
만약 해당 아이디와 비밀번호가 유출된다면 더 큰 피해로 번질 수 있다는 위험이 있다.

때문에, 사용자의 민감한 정보를 서비스에서 직접 저장하고 관리해야한다는 부담감이 생길 것이다. 또한, 네이버와 카카오 같은 서비스에서는 자신의 사용자의 정보를 다른 서비스, 제 3자에게 맡긴다는 것이 불만족스러울 수 있다.

이러한 문제를 해결하기 위해 OAuth가 등장하기 전, 구글은 AuthSub, 야후는 BBAuth처럼 각 회사가 개발한 방법을 사용했다. 하지만 표준화되어 있지 않아 유지보수가 어렵다는 단점이 있었다.

이를 위해 등장한 것이 바로 OAuth이다

## 2. OAuth란?
네이버, 카카오와 같은 플랫폼의 특정한 사용자의 데이터에 접근하기 위해 제3자 클랑이언트(우리의 서비스)가 접근 권한을 위임(Delegated Authorization) 받을 수 있는 표준 프로토콜이다.

즉, 우리의 서비스가 우리의 서비스를 이용하는 사용자의 타사 플랫폼 정보를 얻기 위해 권한을 타사 플랫폼으로부터 위임받는 것이다.

## 3. OAuth 2.0 주체
### Resource Owner
우리의 서비스를 사용하는 사용자.

### Authorization & Resource Server
Authorization Server는 리소스 오너를 인증하고, 클라이언트에게 엑세스 토큰을 발급해주는 서버이다.

Resource Server는 네이버, 카카오와 같이 리소스를 가지고 있는 서버이다.

### Client
Resource Server의 자원을 이용하려고 하는 서비스, 즉 우리의 서비스.

## 4. 어플리케이션 등록
OAuth 2.0을 이용하기 위해서 클라이언트는 리소스 서버에 Redirect URI를 등록해야하는 작업을 거쳐야 한다.

### Redirect URI
Redirect URI는 사용자가 OAuth 2.0 서비스에서 인증을 마치고 (ex. 네이버 로그인 페이지에서 로그인을 마쳤을 때) 사용자를 리디렉션시킬 위치이다.

OAuth 2.0 서비스는 서비스 인증에 성공한 사용자를 사전에 등록된 Redirect URI 로만 리디렉션을 시킨다. 승인되지 않은 Redirect URI로 리디렉션될 경우, Authorization Code를 중간에 탈취 당할 위험이 있기 때문이다.

일부 OAuth 2.0 서비스는 여러 개의 Redirect URI를 등록할 수 있다.

Redirect URI는 기본적으로 보안을 위해 https만 허용되며 localhost를 예외적으로 http를 허용한다.

### Client Id, Client Secret
Redirect UIR 등록 과정 후, Client ID와 Client Secret 을 얻을 수 있다. Client Id와 Client Secret는 액세스 토큰을 발급받는 데에 사용된다.

Client ID는 공개되어도 상관 없지만, Client Secret은 절대 유출되어서는 안된다.

## 5. OAuth 2.0의 동작 매커니즘
- 로그인 요청

리소스 오너가 우리의 서비스에서 네이버로 로그인하기 와 같은 버튼을 클릭해 로그인을 요청한다.

클라이언트는 OAuth 프로세스를 시작하기 위해 사용자의 브라우저는 Authorization Server로 보낸다.
이때, `response_type`, `client_id`, `redirect_uri`, `scope` 등의 매개변수를 쿼리 스트링으로 포함하여 보낸다.

> 매개 변수
> - response_type : code로 값을 설정, 인증이 성공할 경우, Authorization Code를 받을 수 있다.
> - client_id : 어플리케이션을 생성했을 때 발급받은 client id
> - redirect_uri: 어플리케이션을 생성했을 때 등록한 redirect URI
> - scope: 클라이언트가 부여받은 리소스 접근 권한

```http request
https://authorization-server.com/auth?response_type=code
&client_id=123412341234
&redirect_uri=https://musicla-app.com
&scope=create+delete
```

- 로그인 페이지 제공 및 ID/PW 제공

클라이언트가 빌드한 Authorization URL로 이동된 리소스 오너는 제공된 로그인 페이지에서 ID와 PW를 입력하여 인증

- Authorization Code 발급 및 Redirect URI로 리디렉트

인증이 성공되면, Authorization Server는 제공된 Redirect URI에 Authorization Code를 포함시켜 사용자는 리디렉션시킨다.

Authorization Code란 클라이언트가 Access Token을 획득하기 위해 사용하는 임시 코드이다. 이 코드는 수명이 매우 짧다. (일반적으로 1~10분)

- Authorization Code와 Accesss Token 교환

클라이언트는 전달 받은 Authorizaton Code를 Authorization Server에 전달하고, Access Token으로 응답 받는다.

클라이언트는 응답받은 액세스 토큰을 리소스 오너의 액세스 토큰에 저장하고, 이후 Resource Server에 리소스 오너의 리소스에 접근하기 위해 액세스 토큰을 사용한다.

> 매개변수
> - grant_type : 항상 authorization_code 로 설정되어야 한다. (참고)
> - code : 발급받은 Authorization Code
> - redirect_uri : Redirect URI
> - client_id : Client ID
> - client_secret : RFC 표준상 필수는 아니지만, Client Secret이 발급된 경우 포함하여 요청해야한다.

```http request
POST /oauth/token HTTP/1.1
Host: authorization-server.com

grant_type=authorization_code
&code=xxxxxxxxxxx
&redirect_uri=https://example-app.com/redirect
&client_id=xxxxxxxxxx
&client_secret=xxxxxxxxxx
```

- 로그인 성공

위 과정을 성공적으로 마치면, 클라이언트는 리소스 오너에게 로그인 성공을 알린다.

## 6. Authorization Code는 왜 필요할까?
클라이언트는 바로 액세스 토큰을 받아도 될텐데 왜 Authorization code가 필요한 것일까?

Redirect URI를 통해 Authorization Code를 발급받는 과정이 생략된다면, Authorization Server는 Redirect URI를 통해 액세스 토큰을 전달해야 한다. 
Redirect URI를 통해 데이터를 전달하는 방법은 URL 자체에 데이터를 실어 전달하는 방법만 존재하는데, 이렇게 진행하면 데이터가 곧바로 노출된다.

하지만 액세스 토큰은 절대 노출되어서는 안되기 때문에 보안 사고를 방지하가 위해 Authorization Code를 사용하는 것이다.

