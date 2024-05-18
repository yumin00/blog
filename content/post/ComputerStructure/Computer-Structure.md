---
title: "컴퓨터 구조에 대해 알아보자"
date: 2024-05-05T21:20:00+09:00
draft: false
categories :
- ComputerStructure
---


컴퓨터에는 4가지의 핵심 부품이 있다. 그것은 바로 CPU, 주기억장치(메모리), 보조기억장치, 입출력장치이다.
오늘은 컴퓨터의 핵심 부품 중 CPU, 메모리, 보조기억장치에 대해 공부해보고자 한다.

## 컴퓨터 구조
![image](https://github.com/yumin00/blog/assets/130362583/7f45a384-756e-4310-97e2-85c57d53db30)

먼저 컴퓨터의 구조에 대해 간단히 그림으로 살펴보자.
컴퓨터 구조는 사진과 같이 CPU, 메모리, 보조기억장치, 입출력장치고 구성되어 있으며 각 부품들이 서로 통신할 수 있는 시스템버스로 이루어져 있다.
이제 각 장치에 대해 더 자세히 알아보자.

## 시스템 버스
각 부품들이 서로 정보를 주고 받을 수 있는 통로가 바로 시스템 버스이다. 시스템 버스는 주소 버스, 데이터 버스, 제어 버스가 있다.
주소 버스를 통해 서로 주소 정보를 주고 받을 수 있고, 데이터 버스를 통해 서로 데이터를 주고 받을 수 있다.
제어 버스를 통해 서로 제어 신호를 주고 받을 수 있다.

## CPU (중앙처리장치)
![image](https://github.com/yumin00/blog/assets/130362583/d9156833-525c-448f-8154-dcd6ffa1c23d)

CPU는 컴퓨터의 뇌이다. 메모리에 저장된 명령어와 데이터를 읽어와 명령어를 해석/처리하고 명령어의 수행 순서를 제어하는 역할을 한다. 따라서 CPU에는 명령을 처리할 수 있는 장치들을 포함하고 있다.

### ALU (산술논리연산장치)
ALU 는 산술적인 연산과 논리적인 연산을 담당한다.
ALU 는 다음과 같은 과정을 통해 연산 및 결과값을 도출한다.

1. 연산을 위해서는 피연산자와 수행할 연산이 필요한데, 캐시나 메모리로부터 데이터를 읽어오는 데이터 레지스터를 통해 피연산자를 받는다.
2. 제어장치로부터 수행할 연산을 알려주는 제어신호를 받아 연산을 수행한다.
3. 연산 후 결과값은 메모리에 저장하지 않고, 일시적으로 레지스터에 저장한다.
4. 연산 결과에 대한 추가적인 상태 정보(플래그)는 플래그 제시스터로 전송한다.

## 제어장치
제어장치는 제어 신호를 통해 명령어를 해석하고 조작을 지시하는 장치이다. 제어신호는 컴퓨터 부품들을 작동시키기 위한 일종의 전기 신호이다.
CPU가 해석해야 하는 명령어는 명령어 레지스터에 저장되는데, 이때 제어장치는 명령어 레지스터로부터 명령어를 받아 해석한 후 제어 신호를 발생하여 수행해야 하는 내용을 전달한다.
혹은 위 예시처럼, 명령어를 해석한 뒤, 계산해야하는 명령어가 있다면, 수행할 연산을 알려주는 제어신호를 발생시켜 ALU가 연산을 수행할 수 있도록 한다.

### 레지스터
레지스터는 CPU가 요청을 처리하는 데 필요한 데이터를 일시적으로 저장하는 기억장치이다.

1. 프로그래밍 카운터

메모리에서 가져올 명령어의 주소를 저장한다.
2. 명령어 레지스터

메모리에서 가져온 명령어를 저장한다.

3.메모리 주소 레지스터

메모리의 주소를 저장한다. CPU는 메모리 주소 레지스터를 거쳐 읽고자 하는 주소값으로 주소버스로 보낸다.
4. 메모리 버퍼 레지스터(메모리 데이터 레지스터)

메모리와 주고 받은 데이터나 명령어를 저장한다. CPU가 메모리 주소 레지스터를 거쳐 주소버스로 값을 보내면, 해당 데이터는 메모리 버퍼 레지스터에 저장된다.
5. 범용 레지스터

다양하고 일반적인 상황에 자유롭게 사용한다.
6. 플래그 레지스터

ALU 연산 결과에 따른 부가적인 정보(플래그) 저장한다.


### CPU 동작 방식
예를 들어, 메모리에 `주소1과 주소2를 더해라` 라는 명령어가 있다고 가정해보자. 해당 명령어는 어떻게 처리될까?

1. 제어장치는 `메모리 읽기` 제어 신호를 통해, 메모리 주소 레지스터를 거쳐 읽고자 하는 주소값으로 주소버스를 보낸다.
2. 해당 명령어는 메모리 데이터 레지스터에 저장된다.
3. 제어장치는 해당 명령어를 해석하여 주소1과 주소2의 데이터가 필요하다고 판단하여, 다시 `메모리 읽기` 제어 신호를 통해, 메모리 주소 레지스터를 거쳐 읽고자 하는 주소값으로 주소버스를 보낸다.
4. 해당 데이터는 각각 메모리 데이터 레지스터에 저장된다.
5. 제어장치는 ALU에게 연산 제어 신호를 보내고, ALU는 이 신호를 받아 연산을 처리한다.
6. ALU는 결과값을 메모리 데이터 레지스터에 저장한다.
7. 제어장치는 다음 명령어를 읽기 위해 `메모리 읽기` 제어 신호를 통해, 메모리 주소 레지스터를 거쳐 읽고자 하는 주소값으로 주소버스를 보낸다.
8. 해당 명령어를 메모리 데이터 레지스터에 저장된다.
9. 제어장치는 해당 명령어를 해석하고 `메모리 쓰기` 제어 신호와 함께 메모리 데이터 레지스터에 있는 결과값을 함께 전달한다.

## 주기억장치 (메모리)
메모리는 현재 실행되어야 하는 명령어와 데이터를 저장하는 부품이다. 메모리에 있는 데이터를 읽기 위해 주소라는 개념을 사용한다. 필요한 데이터의 주소를 통해 데이터를 얻을 수 있다.

## 보조기억장치
![image](https://github.com/yumin00/blog/assets/130362583/c0f291d0-d0a7-4763-9ca7-8f0130ebd5a0)

위 사진과 같이 4개의 메모리를 비교해볼 수 있다. 보조기억장치에 비해 주기억장치는 비싸고 속도가 느리다는 단점이 있다. 또한, 컴퓨터의 전원을 끄면 주기억장치의 데이터는 모두 휘발된다.
이를 보조하고자, 주기억장치보다 크기가 크고, 전원이 꺼져도 저장된 내용을 잃지 않은 보조기억장치가 생겨나게 되었다.
하드 디스크, SSD, USB 메모리, DVD 와 같은 저장 장치가 보조기억장치이다.
