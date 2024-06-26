---
title: "프로세스의 주소 공간에 대해 알아보자(Process Address Space)"
date: 2024-05-08T17:38:31+09:00
draft: false
categories :
- OS
---

프로세스가 메모리를 할당 받으면, 메모리를 관리하기 위해 이 공간들을 어떤 구조로 관리하는데, 이를 프로세스 주소 공간이라고 한다.
메모리는 한정되어 있기 때문에, 프로세스는 메모리를 절약하기 위해 다양한 방법을 시도하는데, 오늘은 프로세스 주소 공간과 메모리 절약을 위한 방법에 대해 공부해보고자 한다.

## 프로세스 주소 공간
프로세스가 메모리에 할당되면 프로세스마다 고유의 주소 공간이 생기게 된다. 일반적으로 프로세스 주소 공간은 [프로세스 포스팅](https://yumin.dev/p/process-vs-program-vs-thread/)에서 공부한 것처럼 다음과 같이 구성된다.

![image](https://github.com/yumin00/blog/assets/130362583/12ccc8ca-a517-4d87-8673-05e8a7bb42c9)

### Stack
함수 호출과 관계되는 지역변수 혹은 매개변수가 저장되며, 함수 호출이 완료되면 해제되는 영역이다.
메모리의 높은 주소에서 낮은 주소의 방향으로 할당된다.
무한 반복이나 재귀로 인해 스택이 초과되어 에러가 발생하는 상황을 `Stack Overflow` 라고 한다.

### Heap
동적 메모리 할당을 위한 영역으로, 런타임에 크기가 결정되는 영역이다. 주로 참조형 데이터가 할당된다.
메모리의 낮은 주소에서 높은 주소의 방향으로 할당되며, 일반적으로 힙은 실행 중에 계속 확장될 수 있다.

### Data
전역 변수와 정적 변수가 저장되는 영역이다. 실행 파일이 로드될 때, Data 영역에 할당되면 실행이 끝날 때까지 유지된다.

### Code
프로그램의 실제 명령어 코드가 저장되는 영역이다. 프로그램이 실행될 수 있도록 CPU가 해석 가능한 기계어 코드가 저장되어 있는 공간이다.
해당 코드는 수정되면 안 되기 때문에 ReadOnly 상태로 저장되어 있다.
