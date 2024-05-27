---
title: "AWS EC2에 대해 알아보자"
date: 2024-05-13T23:08:57+09:00
draft: true
categories :
- AWS
---

현재 진행하고 있는 프로젝트를 실제로 배포하기 위하여, AWS의 EC2에 대해 공부해보고자 한다.

# EC2란?
AWS(Amazon Web Services)는 전세계적으로 사용되고 있는 클라우트 컴퓨팅 플랫폼이며 다양한 서비스와 기능을 제공하고 있다. 여기에서 클라우드 컴퓨팅이란, 인터넷(클라우드)을 통해서 서버, 스토리지, 데이터베이스 등의 컴퓨팅 서비스를 제공하여 AWS에서 원격으로 제어할 수 있는 가상의 컴퓨터를 한 대 빌리는 것이다.

그 중 EC2(Elastic Compute Cloud)는 사용자가 가상 서버를 쉽게 생성하고 관리할 수 있게 해주는 서비스이다. EC2를 사용하면, 필요한 운영 체제를 선택하고, CPU, 메모리, 저장 공간 등의 사양을 직접 설정하여 관리할 수 있다.

EC2의 장점은 다음과 같다.

- 유연성: 사용자는 필요에 따라 인스턴스의 사양을 선택하고, 언제든지 사양을 변경할 수 있다.
- 확장성: 트래픽이 증가하거나 컴퓨팅 요구사항이 변할 때, 추가 인스턴스를 쉽게 생성하거나 기존 인스턴스를 강화할 수 있다.
- 비용 효율성: 사용한 만큼만 비용을 지불하는 'Pay as you go' 시스템을 통해 비용을 절감할 수 있다.
- 보안: AWS의 보안 인프라를 이용하여 데이터와 애플리케이션을 안전하게 보호할 수 있다.

# EC2 인스턴스 생성 방법에 대해 알아보자
현재 진행하고 있는 프로젝트를 배포하기 위하여, 이번에 직접 인스턴스를 생성해보자!
AWS EC2 인스턴스를 생성한다는 것은 AMI를 토대로 운영체제, CPU, RAM 혹은 런타임 등이 구성된 컴퓨터를 빌리는 것이다. AMI에 대한 더 자세한 내용은 아래에서 확인해보자.

## 1. AWS 접속
<img width="2234" alt="image" src="https://github.com/yumin00/blog/assets/130362583/bca4acf8-650b-42cd-9cd6-53e7a2108799">

먼저, AWS에서 로그인을 하면 위와 같은 화면이 나온다.

<img width="903" alt="image" src="https://github.com/yumin00/blog/assets/130362583/e5f3be59-f178-444f-8d9e-ff2d96bcb6e5">

왼쪽 위의 `서비스` 메뉴를 클릭하여 `EC2 대시보드`로 이동할 수 있다.

## 2. 인스턴스 시작
<img width="929" alt="image" src="https://github.com/yumin00/blog/assets/130362583/77ea7467-8021-456a-8bb5-52c9abd34f88">
EC2 대시보드에서 `인스턴스 시작` 메뉴를 찾아서 `인스턴스 시작` 서비스를 클릭하면 인스턴스를 생성할 수 있다.

## 3. 인스턴스 생성
<img width="819" alt="image" src="https://github.com/yumin00/blog/assets/130362583/11dd05b6-82b8-4225-899d-df09c7228f4b">

해당 화면에서 여러가지를 설정함으로써 직접 인스턴스를 생성할 수 있다. 여러가지 설정해야하는 것들과 옵션들이 많은데 각각 하나씩 알아보자!

### 이름 및 태그
<img width="788" alt="image" src="https://github.com/yumin00/blog/assets/130362583/9d0cdb50-f209-4157-bde5-9558854d9b57">

이름은, "Name" 태그로 저장되며, 인스턴스의 용도/환경 또는 소유자 등을 나타내는ㄷ 사용할 수 있는 이름을 설정할 수 있다.
예를 들어서, 'development-server', 'production-database' 등과 같이 인스턴스의 용도를 쉽게 식별할 수 있는 이름을 부여할 수 있다.

태그는, 키와 값의 상으로 구성된다. 태그를 통해서 사용자는 리소스를 분류하고, 조직 내에서 여러 부서의 비용을 추적하며, 자동화된 리소스 관리 정책을 적용하는 등 다양한 용도로 사용할 수 있다.
예를 들어, 여러 인스턴스 중에서 특정 프로젝트나 부서에 할당된 인스턴스만을 식별하기 위해 태그를 사용할 수 있다. 현재는 하나의 인스턴스만을 생성할 것이기 때문에 태그의 용도를 제대로 확인할 수는 없을 것 같다..

그래서 이번에 인스턴스를 생성할 때는 이름만 설정해보고자 한다.

### 애플리케이션 및 OS 이미지(Amazon Machine Image)
<img width="798" alt="image" src="https://github.com/yumin00/blog/assets/130362583/745c80af-7dd8-43d4-8b57-14749eafcd79">

AMI(Amazon Machine Image)란, 소프트웨어 구성이 기재된 템플릿이다. 즉, 인스턴스를 시작하는데 필요한 정보를 제공하는 이미지로, 한 AMI로 여러 인스턴스를 생성할 수 있다.

어플리케이션이 특정 OS에서만 작동할 경우, 해당 OS를 포함하는 AMI를 선택해야 하며, 최신 상태의 AMI를 선택하는 것이 중요하다!

기본적으로, Amazon Linux 2는 아마존에서 제공하는 것으로, AWS 서비스와 호환성이 뛰어나고, 보안 업데이트/성능 최적화를 위해 지속적으로 관리된다고 한다. 일반적인 용도에 적합하다고 하여 이번 서버 배포의 AMI는 Amazon Linux 2로 설정하여 진행해보고자 한다.

### 인스턴스 유형
<img width="810" alt="image" src="https://github.com/yumin00/blog/assets/130362583/acbc7d47-9dd1-43c4-aa91-19579f766a07">

EC2 인스턴스 유형은 다양한 사용 사례에 맞게 설계된 여러 가지 인스턴스 유형을 제공한다. 작은 웹 서버나 테스트 서버의 경우에는  t3 인스턴스는 경제적이면서도 충분한 성능을 제공한다고 한다.

따라서 이번에는 t3.small 을 설정해보고자 한다. 

### 키 페어
<img width="781" alt="image" src="https://github.com/yumin00/blog/assets/130362583/3f454d04-d817-4841-a819-dff5826806db">

키 페어는 만들어지는 EC2 인스턴스에 대한 안전한 SSH 접속을 위하여 사용된다.
- 공개 키(Public Key): EC2 인스턴스에 저장된다.
- 개인 키(Private Key): 사용자의 로컬 컴퓨터에 저장된다.

<img width="923" alt="image" src="https://github.com/yumin00/blog/assets/130362583/9057ee2f-90db-49f3-8209-d9875f4b23b4">
키 페어 메뉴에서 키 페어를 생성하여 적용할 수 있다!

### 인스턴스 연결
<img width="1533" alt="image" src="https://github.com/yumin00/blog/assets/130362583/90b71bde-eaa3-4a14-bbda-e138bfcffc3e">
이제 인스턴스를 연결해보자. 만든 인스턴스에서 연결을 눌러 인스턴스를 연결할 수 있다.

<img width="856" alt="image" src="https://github.com/yumin00/blog/assets/130362583/60e8cdf1-2f19-42c3-a5fc-acbdc786c0a5">
EC2 인스턴스 연결을 통해 AWS에서 제공하는 터미널 인터페이스를 통해 연결하는 방법이 있고,

<img width="857" alt="image" src="https://github.com/yumin00/blog/assets/130362583/ee2757fa-984b-485f-82c6-1266db526c09">
직접 로컬 터미널에서 EC2 인스턴스를 연결하는 방법이 있다. 나느 AWS에서 제공하는 터미널 인터페이스를 사용하여 진행해보았다.

<img width="1798" alt="image" src="https://github.com/yumin00/blog/assets/130362583/dcc5a8bf-24f4-4957-99d6-dc4951568697">
그러면 이렇게 인스턴스에 연결한 것을 확인할 수 있다.

### 배포
- sudo su
- lsof 0u
- 80이 없음. 80을 실행시켜줘야함 -> git 연결
- 깃헙에서 personala ccess token 발급
- yum install git
- git clone https://${GITHUB_TOKEN}:@github.com/${GITHUB_REPOSITORY}
- home에 yumin 생성 mkdir yumin
- https://kdev.ing/install-docker-compose-in-amazon-linux-2023/ 도커 컴포즈 설치
- sudo service docker start
- docker compose up -d