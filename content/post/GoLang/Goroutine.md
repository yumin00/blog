---
title: "고루틴(Goroutine)에 대해 알아보자"
date: 2023-10-22T21:28:52+09:00
draft: true
categories :
- GoLang
---

# 고루틴(Goroutine)
## 고루틴
고루틴은 go에서 사용하는 경량 스레드를 의미한다. 

go에서는 스레드보다 훨씬 가벼운 비동기 동시 처리를 구현해 각각의 일에 대해 스레드와 1대1로 대응하지 않고, 훨씬 적은 스레드를 사용한다.

go에서는 고루틴을 선언함으로써 함수를 비동기적으로 동시에 실행할 수 있다.

## 고루틴 예제
고루틴을 사용한 함수는 다른 함수와 상관없이 동시에 실행된다.

```go
package main

import "fmt"

func hello() {
	fmt.Println("hello yumin!")
}

func main() {
	go hello()
}
```

해당 함수를 실행시키면 `fmt.Println("hello yumin!")` 는 실행되지 않고 프로그램이 종료된다.
hello() 함수를 고루틴으로 실행시킴으로써 main() 함수와 동시에 실행되기 때문에 hello() 함수의 print가 호출되기 전에 main이 종료되고, 프로그램이 종료된다.
따라서 print가 실행될 때까지 main 함수가 종료되지 않게 대기시키기 위해서 `fmt.Scanln()`를 입력해주어야 한다.

```go
package main

import "fmt"

func hello() {
	fmt.Println("hello yumin!")
}

func main() {
	go hello()
	fmt.Scanln()
}
```

## sync의 WaitGroup
main 함수가 종료되지 않고 대기시키기 위해 `fmt.Scanln()`를 입력해주는 것은 사실 좋은 방법이라고 할 수 있다.

고루틴이 끝날 때까지 main을 대기하는 기능은 바로 sync의 WaitGroup이다. 이는 패키지에 선언되어 있는 고루틴이 모두 종료할 때까지 대기한다.

### sync

- Add() : 기다릴 고루틴의 수 설정
- Done() : 고루틴이 실행된 함수 내에서 호출함으로써 함수 호출이 완료됐음을 알림
- Wait() : 고루틴이 모두 끝날 때까지 차단

### defer 예약어
defer문은 자신을 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행을 연기한다.
    

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func hello(n int, w *sync.WaitGroup){
	r := rand.Intn(3) // 3개의 난수 생성

	time.Sleep(time.Duration(r) * time.Second)

	fmt.Println(n)
    
	defer w.Done() //끝났음을 전달
}

func main() {
	wait := new(sync.WaitGroup) // waitgroup 생성 - new 키워드로 선언한 변수는 포인터형이다.  

	wait.Add(100) // 100개의 고루틴을 기다림. 기다릴 고루틴 개수 설정

	for i := 0; i < 100; i++ {
		go hello(i, wait) // wait을 매개변수로 전달
	}
	wait.Wait() // 고루틴이 모두 끝날때까지 대기 

}
```