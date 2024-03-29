---
title: "Process vs Program vs Thread"
date: 2023-08-29T13:27:13+09:00
draft: false
categories :
- ComputerStructure
---

IPC에 대해 공부하던 중 프로세스란 실행 중인 프로그램을 말한다고 하였는데 해당 개념에 대해 좀 더 자세히 알고 싶어 프로세스와 프로그램의 차이점에 대해 공부해보고자 한다.

# Process vs Program
## Program
프로그램이란, 실행 파일을 의미한다. 파일 시스템에 존재하는 실행 파일이 프로그램이다. 즉, .exe 로 끝나는 파일들이 프로그램인 것이다.

프로그램은 `정적인 상태의 파일`로 존재한다.

## Process
프로세스는 운영체제가 메모리 등의 필요한 자원을 할당해준 `'실행중인 프로그램'`이다. 프로그램을 실행하면 운영체제로부터 실행에 필요한 자원을 할당받아 '프로세스'가 되는 것이다.

프로그램은 하나지만, 이 프로그램을 실행하는 인스턴스는 여러 개가 생길 수 있다. 즉, 프로그램은 하나더라도 프로세스는 여러 개일 수 있다.

Mac에서는 `Activity Monitor.app`을 통해 프로세스를 확인할 수 있다.

<img width="1072" alt="image" src="https://github.com/yumin00/blog/assets/130362583/3096b8ce-ec02-4705-bf21-7b5f3161bdf6">

# Process vs Thread
## Process
[단일 프로세스]

<img width="346" alt="image" src="https://github.com/yumin00/blog/assets/130362583/5fb6dd66-f5b1-47bc-9ba7-f089a3f2997d">

[멀티 프로세스]

<img width="708" alt="image" src="https://github.com/yumin00/blog/assets/130362583/5854eea9-5387-476c-8cc8-2fd98275e63d">

프로세스는 Code, Data, Stack, Heap의 구조로 되어있는 독립된 메모리 영역을 할당 받는다. 또한, 프로세는 최소 한 개 이상의 스레드를 갖는다.

각 프로세스는 별도의 주소 공간에 할당되고, 서로 독자적인 메모리를 갖기 때문에 서로 메모리를 공유할 수 없다.

각 프로세스가 서로 통신하기 위해서는 프로세스 간 통신(IPC)를 사용해야 한다.

- 프로세서 : CPU
- 멀티 프로세싱 : 여러 개의 프로세스를 사용
- 멀티 태스킹 : 같은 시간에 여러 개의 프로그램을 띄우는 것

여러 프로세스를 한번에 돌리는 작업은 동시, 병렬성 또는 이 둘의 혼합으로 이뤄진다.

### 동시성
프로세서는 원래 하나의 프로세스만 실행시킬 수 있다. 동시성은 프로세서 하나가 여러개의 프로그램을 돌아가면서 일부분씩 수행하는 방식이다.

진행 중인 작업을 바꾸는 것을 context switching 이라고 한다. 이 과정이 매우 빠른 속도로 진행되기 때문에, 사람들에게는 여러개의 프로그램이 동시에 실행되는 것처럼 보인다.

### 병렬성
프로세서 하나에 여러개의 코어가 달려서 여러 개의 프로그램을 각각 동시에 작업하는 방식이다.

---

브라우저 또한 하나의 프로그램이고, 브라우저가 도는 동안 하나의 프로세스가 실행된다. 하지만 브라우저가 실행중일 때도, 노래를 듣는 등 한 프로세스 안에서도 여러가지 작업들이 동시에 진행되는 경우도 있다.

이러한 여러가지 작업으로 생기는 이 갈래들을 스레드(Thread)라고 한다.

## Thread
[단일 스레드]

<img width="262" alt="image" src="https://github.com/yumin00/blog/assets/130362583/286eb377-cf11-4639-806f-1ff309363f63">

[멀티 스레드]

<img width="692" alt="image" src="https://github.com/yumin00/blog/assets/130362583/989a5cdf-01fd-4751-b365-ac795eefddbd">

스레드(thread)는 프로세스가 할당 받은 자원을 이용하여 실제로 작업을 수행하며, 프로세스의 특정한 수행 경로이자 프로세스 내에서 실행되는 여러 흐름의 단위이다.

프로세스가 운영체제로부터 자원을 할당 받으면 그 자원을 스레드가 사용한다. 프로세스는 최소 한 개 이상의 스레드를 가지며, 이 스레드를 메인 스레드라고 한다.

스레드는 프로세스가 할당받은 메모리에서 독자적인 stack 영역을 갖는다. 여러 스레드끼리는 메모리를 공유하기 때문에, 동기화/데드락 등의 문제가 발생할 수 있다. 

## 멀티 프로세스 vs 멀티 스레드
### 멀티 프로세스
- 하나의 응용 프로그램을 여러 개의 프로세스로 구성하여 각 프로세스가 하나의 테스크를 처리하는 것
- 여러 개의 자식 프로세스 중 하나에 문제가 발생하면 그 자식 프로세스만 죽는다.
- Context Switching 과정에서 캐쉬 메모리 초기화 등 무거운 작업이 진행되고 많은 시간이 소모되는 등의 오버헤드 발생

> 오버헤드
> 
> 특정 작업을 위해 사용되는 간접적인 처리 시간, 메모리 등


### 멀티 스레드
- 하나의 응용프로그램을 여러 개의 스레드로 구성하고 각 스레드가 하나의 작업을 처리
- 시스템 콜이 줄어들어 자원 관리가 용이
- 하나의 스레드에 문제가 발생하면 전체 프로세스가 영향을 받음
- 자원 공유에서 동기화 문제가 발생할 수 있음

프로세스/스레드 간의 통신은 IPC를 통해 이뤄질 수 있는데, 더 자세한 내용은 [여기](https://yumin.dev/p/ipc/)에서 볼 수 있다.