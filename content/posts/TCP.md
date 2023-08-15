---
title: "TCP"
date: 2023-08-15T21:25:00+09:00
draft: true
---

# TCP
## 1. TCP
### TCP란?
TCP(Transmission Control Protocol)이란, 신뢰성 있는 데이터 통신을 위한 프로토콜이다.

IP는 오직 데이터를 주고받을 주소만을 위한 프로토콜이지만, TCP는 주소들이 주고받을 데이터가 온전한지를 확인해주는 프로토콜이다.

TCP의 패킷 유실 문제 해결 방법
1. TCP는 어떤 패킷이 사라졌는지 확인하기 위해 패킷마다 번호를 붙인다.
2. TCP는 클라이언트가 데이터를 받으면, 서버에게 데이터를 잘 받았다는 메세지를 전송한다.
3. 서버가 클라이언트에게 데이터를 잘 받았다는 메세지를 받지 못하면, 데이터 전송 과정에 오류가 생겼음을 인지하고 다시 데이터를 보내, 클라이언트가 데이터를 받을 수 있게 보장한다.

### 헤더 / 플래그
IP를 포함해 모든 프로토콜은 데이터의 앞에 헤더라는 정보를 추가해 전송한다. 헤더 안에는 해당 데이터에 대한 정보가 담겨 있다.(데이터가 들어있는 것이 아니라, 데이터의 정보가 들어있음)

TCP는 데이터의 신뢰를 담당하는 프로토콜이기 때문에, 헤더에 많은 정보가 들어가있다. 그중에서도 패킷의 상태를 알리는 목적의 헤더 정보인 플래그가 있다.

플래그는 다른 기기에 신호를 전달하기 위한 용도로, 플래그의 활성화 여부에 따라 다른 의미를 전달할 수 있다.

[플래그의 종류]
- ACK(Acknowledgement)
  - 앞서 받은 데이터를 잘 처리했다는 의미를 가지는 플래그
- SYN(Synchronize)
  - 연결을 요청하는 플래그
- FIN(Finish)
  - 통신이 마무리되어 연결의 해제를 요청하는 플래그

> (예시)
>
> 플래그 : ACK : 1 / SYN : 1
>
> =
> 상대방이 보낸 데이터를 잘 처리했다(ACK) / 나도 상대방에게 연결해도 되는가?(SYN)

## 2. 연결형 프로토콜, TCP
TCP는 신뢰성 있는 통신을 위한 프로토콜로, 데이터를 주고받기 전에 미리 클라이언트와 서버가 서로 통신할 준비가 됐는지 확인하고, 연결 오류로 이한 데이터 유실을 방지한다. -- 연결형 프로토콜

### 3방향 핸드셰이크
TCP는 3방향 핸드셰이크 방식을 사용해 연결을 수립한다.

ex) 클라언트가 서버에게 고양이 사진을 받고 싶다고 요청했다
1. 클라이언트는 서버에 연결하고 싶다는 패킷 전송 >> 플래그 SYN = 1
- SYN이 켜져있어야 서버는 이 패킷이 연결을 시작하고 싶다는 요청이라는 것을 알 수 있다.
2. 클라이언트는 자기의 상태를 연결을 요청한 상태(SYN_SENT)로 바꿔 서버의 답장을 기다림
3. 서버는 클라이언트의 패킷을 확인했다는 답장과 동시에 자신도 클라이언트에서 연결하고 싶다는 요청의 패킷을 전달 >> 플래그 ACK = 1, SYN = 1
4. 서버는 클라이언트의 연결 요청을 받은 상태(SYN_RECEIVED)로  바뀜
5. 클라이언트는 ACK 플래그를 보고 연결해도 된다고 생각해, 자신의 상태를 연결됨(ESTABLISHED)으로 변경하고, SYN에 대한 답장으로 자신 또한 서버에 ACK 메세지를 전달
6. 서버는 답장을 확인한 뒤, 자신의 상태를 연결되었다고 바꿈으로써 연결에 성공함

==> 3방향 핸드쉐이크는 클라이언트와 서버가 각각 연결 요청을 보내고, 서로 확인했다는 응답을 받으면 연결에 성공하는 방식으롣 동작한다.

### 4방향 핸드쉐이크
TCP는 더이상 보낼 데이터가 없는데 연결을 지속하는 곳은 비효율적으로 생각해 데이터 통신이 끝나면 연결을 끊는 과정을 거친다. 이러한 연결 종료 방식을 4방향 핸드셰이크라고 한다.

1. 필요한 데이터를 전부 받았다고 생각한 클라이언트는 서버에게 FIN 플래그가 담긴 패킷을 보내고, 종료를 기다리는 상태(FIN_WAIT)로 변경한다.
2. 연결 종료 요청을 받은 서버는 확인했다는 ACK 메세지를 보낸다.
3. 그리고 미처 보내지 못한 패킷을 마저 보내며 자신의 통신이 끝날 때까지 기다리는 상태(CLOSE_WAIT)가 된다.
4. 시간이 지난 뒤, 서버가 자신 또한 연결을 종료할 준비가 되었다면 클라이언트에게 FIN 플래그를 전송하고 ACK 플래그를 기다리는 상태(LAST_ACK)가 된다.
5. 마지막으로 클라이언트는 서버에게 ACK 플래그가 담긴 패킷을 보내 연결을 종료한다.(TIME_WAIT)
6. 해당 플래그를 받은 서버는 종료한다.(CLOSED)

3방향 핸드셰이크와 달리, 클라이언트와 서버의 마지막 상태가 서로 다른데, 그 이유는 클라이언트는 미처 받지 못한 패킷이 들어오는 등 혹시 모를 상황에 대비해 잠시 기다렸다가(TIME_WAIT) 지정된 시간이 지나면 연결을 종료하는 상태(CLOSED)로 바뀌기 때문이다.

TCP는 서로 연결하고 해지하는 과정을 거치면서 안전하게 데이터를 주고 받는다.