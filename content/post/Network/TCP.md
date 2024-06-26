---
title: "TCP(Transmission Control Protocol)에 대해 알아보자"
date: 2024-04-01T19:44:00+09:00
draft: false
categories :
- Network
---

# TCP에 대해 알아보자
TCP에 대해 공부하기 전, IP 특징에 대해 다시 한번 살펴보자.

IP는 패킷을 받을 대상이 연결 상태인지 확인하지 않고 전송한다. 또한 순서가 있는 여러 개의 패킷을 전송할 때는 순서를 보장할 수 없다. 즉 **비연결성**과 **비신뢰성** 이라는 특징을 가진다.
(IP에 대한 더 자세한 설명은 [여기](https://yumin.dev/p/ip%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90/)에서 확인할 수 있다.)

# TCP는 왜 생겨났을까?
IP만으로는 패킷 전달에 여러 한계점이 있어서 TCP가 만들어졌다.

# TCP란?
TCP(Transmission Control Protocol)란, 데이터를 신뢰성 있게 전달할 수 있도록 해주는 프로토콜이다.

데이터를 주고받을 때, TCP 덕분에 데이터가 안정적으로, 순서대로 도착할 수 있는 것이다.

IP는 오직 데이터를 주고받을 **주소**만을 위한 프로토콜이지만, TCP는 주소들이 주고받을 **데이터**가 온전한지를 확인해주는 프로토콜이다.

# TCP 특징
- **신뢰성**: 데이터 분실, 중복, 오류가 발생할 경우 재전송을 수행한다.
- **연결 지향적**: 통신을 시작하기 전, 송수신자 간의 연결을 설정한다.
- **흐름 제어**: 네트워크의 혼잡도를 감지하고, 데이터 전송 속도를 조절하여 효율적으로 통신을 유지한다.
- **순서 보장**: 데이터가 순서대로 목적지에 도착한다.

# TCP의 패킷 유실 문제 해결 방법
TCP는 어떤 패킷이 사라졌는지 확인하기 위해서 패킷마다 번호를 붙인다. 또한 클라이언트가 데이터를 받으면, 서버에게 데이터를 잘 받았다는 메시지를 전송한다.

만약, 서버가 클라이언트에게 데이터를 잘 받았다는 메세지를 받지 않으면, 데이터 전송 과정에 오류가 발생했음을 인지하고 다시 데이터를 전송하여 클라이언트가 데이터를 받을 수 있도록 한다.

# TCP의 동작 과정
TCP는 3단계를 통해 동작한다. 동작 과정에 대해 알아보기 전, 먼저 플래그에 대해서 알아야 한다.

## 플래그
모든 프로토콜은 데이터 앞에 헤더라는 정보를 추가하여 전송한다. 헤더에는 데이터에 대한 정보가 들어있는데, 헤더엔느 플래그가 포함되어 있다.

플래그는 다른 기기에 신호를 보내기 위한 것인데, 어떤 플래그를 전송하냐에 따라 다른 의미를 전달할 수 있다.

- ACK(Acknoledgement) : 받은 데이터를 잘 처리했다.
- SYN(Synchronize): 연결을 요청한다.
- FIN(Finish): 연결 해제를 요청한다.

## 1. 연결 설정 (3-Way Handshake)
![image](https://github.com/yumin00/blog/assets/130362583/d76ce36f-56f6-4a0f-8799-f1c99790b734)

TCP는 3방향 Handahke 방식을 사용하여 클라이언트와 서버의 연결을 확인한다. 

클라이언트가 서버에게 강아지 사진이 받고 싶다고 요청했다고 가정해보자!

먼저, 클라이언트는 서버에게 연결하고 싶다는 **SYN** 플래그가 담긴 패킷을 전송한다. SYN이 켜져있어야만 서버는 이 패킷이 연결을 시작하고 싶다는 요청임을 알 수 있다.

클라이언트는 자신의 상태를 연결을 요청한 상태(**SYN_SENT**)로 바꾸고 서버의 답장을 기다리게 된다.

서버는 클라이언트의 패킷을 확인했다는 플래그 **ACK**와 자신도 클라이언트에게 연결하고 싶다는 플래그 **SYN**이 담긴 패킷을 전달한다.

그리고 서버의 상태는 연결 요청을 받은 상태(**SYN_RECEIVED**)로 바뀌게 된다.

클라이언트는 ACK 플래그를 통해 연결해도 된다고 판단한여, 자신의 상태를 **ESTABLISHED**(연결됨)으로 변경하고 서버의 SYN 플래그에 대한 답장으로 **ACK** 답장을 전달한다.

서버는 답장을 확인한 뒤, 자신의 상태를 **ESTABLISHED** 로 변경하고, 이를 통해 클라이언트와 서버는 서로 연결되었음을 파악한다.

## 2. 데이터 전송
![image](https://github.com/yumin00/blog/assets/130362583/a4acc645-65b1-45ae-9ce0-a7b9aa6bcbc3)
서버는 클라이언트에게 데이터를 전송하고, 클라이언트는 데이터를 받았다면 ACK를 전송한 뒤, 서버는 클라이언트가 잘 받았음을 확인하게 된다.

## 3. 연결 종료 (4-Way Handshake)
![image](https://github.com/yumin00/blog/assets/130362583/cf8da44b-58ec-4c0f-a05e-0b93a17eef24)
TCP는 더이상 보낼 데이터가 없는데 클라이언트와 서버에 계속 연결하고 있는 것은 비효율적이라고 생각하기 때문에,
데이터 통신이 끝나면 연결을 끊는 과정을 거친다. 해당 방법을 4방향 Handshake ㅌ라고 한다.

필요한 데이터를 전부 받았다고 생각하는 클라이언트는 서버에게 연결 해제를 요청하는 **FIN** 플래그가 담긴 패킷을 전송하고, **FIN_WAIT** 상태가 된다.

연결 종료 요청을 받은 서버는 해당 패캣을 확인했다는 ACK 플래그와 미처 보내지 못한 패킷을 마저 보내며 자신의 통신이 끝날 때까지 기다리는 상태인 **CLOSE_WAIT** 상태가 된디.

시간이 지난 뒤, 서버가 자신 또한 연결을 종료할 준비가 되었다면, 클라이언트에게 **FIN** 플래그가 담긴 패킷을 전송하고, ACK 플래그를 기다리는 **LAST_ACK** 상태가 된다.

클라이언트는 서버에게 **ACK** 플래그가 담긴 패킷을 보내 연결을 종료한다. (**TIME_WAIT**)

해당 패킷을 받은 서버는 종료되고 **CLOSED** 상태가 된다.

4방향 Handshake는 3방향 Handshake와 다르게, 클라이언트와 서버의 마지막 상태가 다르다. 이는 클라이언트는 미처 받지 못한 패킷이 들어오는 등 혹시 모를 상황에 대비해서 잠시 기다렸다가 (TIME_WAIT)
지정된 시간이 지나면 연걸을 종료하는 상태(CLOSED)로 변경되기 때문이다.


# TCP/IP 4계층
TCP로 서로 데이터를 주고 받을 때, 클라이언트와 서버는 3방향 핸드셰이크를 통해 서로 연결된다고 했다. 그러면, 실제로 어떤 과정을 통해 클라이언트가 서버에게 연결을 요청하는걸까?

이러한 내용은 TCP/IP 4계층을 통해 알 수 있다.

## TCP/IP 4계층
![image](https://github.com/yumin00/blog/assets/130362583/5038e7c6-80cd-40ca-95d2-eb75b51b3fb1)
### 1. Application Layer
어플리케이션 계층에서는 사용자와 가장 가까운 계층이다.

사용자가 웹사이트에 접속하려고 한다고 가정해보자. 이때, 사용자는 URL은 입력할 것이고 웹 브라우저는 HTTP 를 통해 웹 서버와 통신하게 될 것이다.
URL에 포함된 도메인 이름은 도메인 이름 시스템(DNS)을 통해 IP 주소로 변환된다. 즉, 클라이언트가 서버에 연결하기 전에, DNS 조회를 통해 서버의 IP 주소를 알아내는 것이다.

어플리케이션 계층에서 사용되는 프로토콜은 HTTP / HTTPS / FTP 등이 있다.

### 2. Transport Layer
전송 계층에서는 TCP인지 UDP인지 결정하여 이를 인터넷 계층에 전달한다.

전송 계층에서는 데이터를 신뢰성 있게 전달하는 역할을 한다.

### 3. Internet Layer
인터넷 계층에서는 보내는 이가 누구이고, 누구에게 보낼 것인지에 대한 내용을 적는다. 그리고 해당 내용을 네트워크 인터페이스 계층에 전달한다.

인터넷 계층에서는 데이터를 목적지까지 전달하는 라우팅을 담당한다.

### 4. Network Interface Layer
IP는 인터넷 주소이기 때문에, 실제 물리적 주소는 알 수 없다. 네트워크 인터페이스 계층에서는 IP를 기반으로 물리적 주소인 MAC를 통해 두 장치 간의 데이터 전송을 담당한다.

## 동작 방식
클라이언트가 정보를 요청하면 각 레이어를 거치면서 캡슐화가 진행된다. 각 레이어들은 추가한 헤더 정보를 덮어서 다음 레이어에게 전달한다.

서버의 네트워크 인터페이스 레이어가 데이터를 받으면 해당 데이터를 역캡슐화하며 데이터를 까면서 상위 레이어에게 전달한다. 그러면 각 레이어들은 데이터를 통해서 어디로 전송해야하는지에 대해 판단하고 다시 상위 레이어에게 전달한다.

사용자가 웹 서버의 html을 요청했다면, 서버는 이 요청을 받고 전달하고자 하는 데이터를 Application Layer부터 데이터를 감싸면서 하위 레이어에게 전달하고,

클라이언트의 네트워크 인터페이스 레이어는 해당 데이터를 받아서 역캡슐화를 진행하며 다시 상위 레이어에게 전달하고, Application Layer가 최종적으로 데이터를 받아서 사용자에게 데이터를 보여주는 것이다.


## TCP의 제어
패킷은 순서대로 전송되고, 순서대로 조립되는 것이 중요하다. 외부상황에 의해 데이터가 유실되거나 순서가 잘못되어 잘못 수신될 수가 있는데 이러한 상황은 언제 발생하고, 어떻게 해결할 수 있을까?

### 흐름제어 (Flow Control)
TCP에서 수신자가 데이터를 받지 못하는 원인 중 하나는, 전송자와 수신자 간에 데이터 속도 차이가 발생했을 경우이다.

이러한 문제를 해결하기 위해서 전송자는 수신자의 처리 속도에 맞춰 통신 속도를 제어해야하는데, 이를 **흐름 제어**라고 한다. 흐름제어하는 방식에는 여러 가지가 있다.

### 정지-대기 방식
정지-대기 방식은 전송한 패킷을 잘 받았다는 응답을 받았을 경우에믄 다음 패킷을 전송하는 것이다.

이 방식은 답장이 없으면 다음 데이터를 전송할 수 없기 떄문에 시간이 오래 걸린다는 단점이 있어 오늘날에는 거의 사용하지 않는다.

### 슬라이딩 윈도 방식
슬라이딩 윈도 방식은, 수신자가 받을 수 있을만큼의 데이터만 보내는 것이다.

![image](https://github.com/yumin00/blog/assets/130362583/8ae45e61-d548-4166-9d95-ba28855d0ccd)
수신자는 받을 수 있는 데이터의 크기를 헤더의 구성요소에 넣어 보내고, 전송자와 수신자가 서로 연결하는 3방향 핸드셰이크 과정에서 윈도 크기가 정해진다.

![image](https://github.com/yumin00/blog/assets/130362583/af044df4-ad11-4c72-86b5-f42fee9033f4)
수신자는 받은 데이터를 처리한 뒤, 응답 패킷을 보낼 때, 마지막으로 받은 데이터가 몇번이었는지 같이 전송하여 전송자는 이를 보고 다음 패킷을 전송하는 방식으로 이루어진다.

## 혼잡제어(Congestion Control)
서버는 데이터를 지역망 도는 인터넷 네트워크를 통해 전달하는데, 만약 특정 네트워크에 데이터가 집중되면 해당 네트워크를 사용하는 데이터의 처리 속도는 저하된다.

하지만 클라이언트는 데이터가 도착하지 않으면 네트워크의 문제인지 파악하지 못한채, 데이터가 오지 않았으니 다시 보내달라고 요청을 한다. 이렇게 되면 클라이언트는 중복으로 데이터를 받는 문제가 발생한다.

네트워크 내에 패킷의 수가 과도하게 증가하는 현상을 혼잡이라고 하며, 이를 제어하는 기능을 **혼잡 제어**라고 한다.

혼잡제어 방식은 처음에는 데이터를 천천히 보내다가, 수신자가 잘 받는 것을 확인하면 점차 속도를 높이고, 수신자가 한참 전에 전달한 데이터까지만 받았다는 응답을 보내면 네트워크가 혼잡하다고 파악하고 속도를 줄이는 것이다.

혼잡제어 방식을 실행하는 방식에는 여러가지가 있다.

### 합 증가/곱 감소
처음에 서버가 클라이언트에게 패킷을 하나씩만 보내고, 응답을 받으면 윈도 크기를 1씩 증가시키면서 전송하는 방식이다. 만약, 네트워크가 혼잡하다고 판단되면 윈도 크기를 절반으로 줄여 네트워크가 안정화되기를 기다린다.

이 방식은 네트워크를 안전하게 사용할 수 있지만, 전송 속도가 천천히 증가하기 때문에 초기 전송 속도를 높이는 데에는 걸리는 시간이 길다.

### 느린 시작
윈도 크기를 1로 시작해 서서히 개수를 늘리는 방식으로, 윈도 크기를 2배씩 증가시킨다. 그리고 네트워크가 혼잡하다가 판단되면 윈도 크기를 1로 급격히 줄인다.

이 방식은 시간이 지날수록 많은 양의 데이터를 전달할 수 있다.

현재 TCP에서는 처음에는 느린 시작을 사용했다가, 특정 지점을 넘기면 합 증가/곱 감소로 변경하는 등 각 상황에 맞는 방법을 조합하여 혼잡 제어 정책을 선택하여 사용하고 있다.
