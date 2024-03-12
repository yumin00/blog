---
title: "Context Switching에 대해 알아보자"
date: 2024-03-12T17:55:26+09:00
draft: false
categories :
- ComputerStructure
---

# Context Switching
## Context란?
context switching에 대해 이야기하기 전, 먼저 context에 대한 이해가 필요하다.

프로세스는 자신만의 PCB(Process Control Block)을 갖는다. PCB에 대한 자세한 이야기는 [여기](https://yumin.dev/p/pcbprocess-control-block%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90/)에서 확인할 수 있다.

[PCB]

![image](https://github.com/yumin00/blog/assets/130362583/9ff52fe0-7789-48a7-8063-f4559a0e0026)

이때, context는 PCB에 저장된다. context는 CPU가 해당 프로세스를 실행하기 위해서 필요한 정보들을 말한다.

## Context Switching 이란?
### context switching 이 필요한 이유는?
CPU는 한 번에 하나의 프로세스만 실행시킬 수 있다.
하지만 OS가 CPU에게 다른 프로세스를 실행시키라고 명령했을 때, 해당 프로세스를 멈추고 다른 프로세스로 변경해야하기 때문에 context switching이 필요하다.

기존의 프로세스의 상태와 context를 저장하고, CPU가 다음 프로세스를 실행시킬 수 있도록 새로운 프로세스의 상태와 context를 교체하는 작업을
**context switching** 이라고 한다.

> 프로세스 상태
>
>- 생성 상태 : PCB를 할당받았지만, CPU의 할당을 기다리는 상태
>- 준비 상태 : CPU를 할당받았지만, 아직 차례가 아닌 상태
>- 실행 상태 : CPU를 할당받아 실행 중인 상태
>- 대기 상태 : 입출력 작업을 요청한 프로세스가 입출력장치의 작업을 기다리는 상태
>- 종료 상태 : PCB와 프로세스가 사용한 메모리는 정리됨

### context switching 은 어떻게 진행될까?

![image](https://github.com/yumin00/blog/assets/130362583/6dbd700e-cd22-4eac-b52f-5440766bd744)

- 현재 실행 중인 프로세스의 정보는 CPU 내부의 레지스터에 저장하고 있다.
- 만약, 시스템 콜로 다른 프로세스를 실행해야한다면, CPU는 현재 실행중인 프로세스의 정보를 PCB에 저장한다. 
- 다음 실행할 프로세스의 PCB 정보를 Regiser에 적재한다.

## Context Switching 을 하는 주체
context switching 은 system call 혹은 interrupt에 의해 발생할 수 있다.

### Interrupt(인터럽트)
인터럽트란, 예외상황이 발생했을 때, CPU에d게 알려 처리할 수 있도록 하는 것을 말한다. 예외상황은 다음과 같은 상황이 있을 수 있다.

### System call
응용 프로그램은 운영체제(OS)가 제공하는 서비스에 접근하기 위해서는 system call 이라는 인터페이스를 사용해야 한다.

만약, 사용자가 크롬 어플리케이션을 사용하고 있다면 운영체제는 CPU에 크롬 프로세스를 할당했을 것이다.
사용자가 카카오톡 어플리케이션으로 전환하려고 한다면, 운용체제는 카카오톡 프로세스의 PCB를 CPU 레지스터에 적재할 것이다.