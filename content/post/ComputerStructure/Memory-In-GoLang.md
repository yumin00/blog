---
title: "GoLang의 메모리 할당에 대해 알아보자"
date: 2023-12-06T23:01:25+09:00
draft: true
categories :
- ComputerStructure
---

# GoLang의 메모리 할당에 대해 알아보자
Go언어에서는 메모리 할당을 어떻게 하고 있을까?

메모리 할당에 앞서 스택과 힙에 대해 복습해보자.

## Stack vs Heap
### Stack
- LIFO(Last In, First Out)의 구조로 이루어져 있다.
- 함수 호출과 함께 자동으로 관리되는 메모리 영역
- 매우 빠르게 할당됨
- 함수 호출이 발생할 때 메모리에 할당되고, 함수가 반환될 때 해제된다.
- 스택의 크기는 제한적이고, 함수 호출에 의해 자동으로 관리되기 때문에 크기가 크거나 수명이 긴 데이터는 적합하지 않다.
- 함수의 매개변수, 반환 주소, 지역 변수

### Heap
- 프로그램이 실행되는 동안 동적으로 할당되는 메모리 영역
- 개발자가 수동으로 메모리를 할당하고 해제해야 됨
- 큰 데이터나 수명이 긴 데이터를 다루는 데 적합함
- 힙 메모리의 할당과 해제는 스택에 비해 느림
- 메모리 누수와 같은 문제가 발생할 수 있음
- 전역 변수

## Go에서의 Stack vs Heap
Go에서 대부분의 메모리 할당은 컴파일러에 의해 자동으로 관리된다.
컴파일러는 변수의 수명와 사용 범위를 분석하여 적절한 메모리 영역에 할당하며, 가비지 컬렉터가 힙 메모리를 자동으로 정리해준다. 이로 인해, 개발자는 메모리 관리에 대한 부담감을 줄일 수 있다.

### Escape Analysis(이스케이프 분석)
Go언어는 정적 타입 언어로 자료형을 지정해서 코드를 작성해야 한다. 따라서, 컴파일할 때 메모리 할당이 모두 결정되는데 이때 사용되는 규칙이 바로 이스케이프 분석이다. 즉, 컴파일러는 이스케이프 분석 기술을 사용하여 스택 할당과 힙 할당 중 어떤 것을 사용할지 선택한다.

이스케이프 분석이란, 객체의 포인터가 서브 루틴 밖으로 전파되는지 분석하는 기술이다.

- 수명이 특정 범위로 한정되거나, 메모리 크기가 컴파일 시에 확정될 경우 스택 할당
- 확정할 수 없는 경우, 힙 할당
- 한 함수 내에서 정의된 변수가 해당 함수 밖으로 빠져나가는 경우 힙 할당
- 하위 함수로의 메모리 전달은 스택을 이용해 메모리 할당 비용을 줄인다
- 상위 함수로의 메모리 전달은 힙에 할당됨
- go build -gcflags "-m" 같이 옵션을 지정하여 빌드하면 탈출 분석 결과가 출력됨

### Stack
- 함수 내 선언된 지역 변수는 기본적으로 스택에 할당 
- 함수 호출이 끝날 때 해당 변수의 수명은 끝남
- 변수가 함수 내에서만 사용되고, 함수 호출이 끝날 때 해당 변수의 수명이 끝난다면 해당 변수는 스택에 할당됨

### Heap
- 이스케이프 분석: 함수 외부로 탈출하는 변수를 자동으로 힙에 할당
- 포인터를 사용하여 변수를 참조할 경우 힙에 할당
- `new`, `make` 같은 함수를 통해 동적으로 메모리를 할당하는 경우, 힙에 할당
- 가비지 컬렉터가 힙 메모리를 자동으로 관리해주기 때문에 개발자가 수동으로 메모리를 해제할 필요가 없음

### Go 메모리 할당
- 힙 할당은 무겁다고 한다. 그 이유는 할당 후 가비지 컬렉터가 할당된 객체 참조 여부를 주기적으로 검사하기 때문이다.
- 스택 할당은 가볍다고 한다. 스택 할당의 경우 CPU 명령어 PUSH/POP 두 개로 끝나기 때문이다.


### TEST
해당 테스트는 go의 1.21 버전을 기준으로 진행해보았다.

먼저 `go build -gcflags "-m"` 옵션을 통해 탈출 분석을 하면, 다양한 결과를 얻을 수 있는데 결과에 대한 뜻을 알아보고자 한다.

- can inline: 스택에 할당됨
- does not escape: 힙에 할당되는 경우라고 생각할 수 있지만, 스택에 할당됨
- escapes to heap: 힙에 할당됨
- moved to heap: 스택에 할당했다가 힙으로 옮기는 것을 의미한다.

```go
package main

type user struct {
	name string
	age  int
}

func main() {
	_ = test()
}

func test() int {
	yumin := &user{
		name: "yumin",
		age:  24,
	}
	return yumin.age
}
```

```
 ./main.go:12:6: can inline test
./main.go:8:6: can inline main
./main.go:9:10: inlining call to test
./main.go:9:10: &user{...} does not escape
./main.go:13:11: &user{...} does not escape
```

`&user{...} does not escape`
해당 테스트에서 yumin은 user의 참조형으로 선언되었지만, 수명이 특정 범위로 한정되기 때문에 스택에 할당되는 것을 확인할 수 있다.

```go
package main

type user struct {
	name string
	age  int
}

func main() {
	_ = test()
}

func test() *user {
	u := user{
		name: "yumin",
		age:  24,
	}

	return &u
}
```

```
./main.go:12:6: can inline escapeToHeap
./main.go:8:6: can inline main
./main.go:9:18: inlining call to escapeToHeap
./main.go:13:2: moved to heap: u
```

- `moved to heap: u`: test()에서 u는 user로 스택에 할당됐지만, u가 포인터형으로 반환되기 때문에 힙으로 이동된 것을 알 수 있다. Go 컴파일러는 기본적으로 가능한 한 변수를 스택에 할당하려고 시도하기 때문에 처음에 스택에 할당되는 것을 알 수 있다.

```go
package main

import "fmt"

func main() {
	x := 10
	y := square(&x)
	fmt.Println(*y)
}

func square(x *int) *int {
	z := (*x) * (*x)
	return &z
}

```

```
./main.go:16:6: can inline square
./main.go:12:13: inlining call to square
./main.go:13:13: inlining call to fmt.Println
./main.go:13:13: ... argument does not escape
./main.go:13:14: *y escapes to heap
./main.go:16:13: x does not escape
./main.go:17:2: moved to heap: z
```

- square()에서 *x는 참조형이지만, 수명이 특정 범위로 한정되기 때문에 스택에 할당되는 것을 확인할 수 있다.
- z는 참조형이 아니기 때문에 `z := (*x) * (*x)` 에서 스택에 할당된다. 하지만 &z 로 포인터형으로 반환되기 때문에 스택에서 힙으로 옮겨지는 것을 확인할 수 있다.
처음에 스택에 할당된 이유는 스택에 할당된 x를 곱하여 만들어진 z를 힙에 바로 할당하는 것보다 스택에 할당하는 것이 속도면에서 더 빠르다는 컴파일러의 판단으로 스택에 할당됐다가 힙으로 옮겨지는 것이라고 판단할 수 있다.(컴파일러는 기본적으로 가능한 한 변수를 스택에 할당하려고 시도한다.)

```go
package main

import "fmt"

func main() {
	var p *int
	fmt.Println(*foo(p))
}

func foo(p *int) *int {
	fmt.Println(p)
	x := 10
	p = &x
	return p
}
```

```
./main.go:11:13: inlining call to fmt.Println
./main.go:7:13: inlining call to fmt.Println
./main.go:10:10: leaking param: p
./main.go:12:2: moved to heap: x
./main.go:11:13: ... argument does not escape
./main.go:7:13: ... argument does not escape
./main.go:7:14: *foo(p) escapes to heap
```

- `x := 10` 에서 x는 스택에 할당된다.
- `moved to heap: x`: 하지만 p가 x의 참조형을 사용하여 반환되기 때문에 x는 힙으로 옮겨지는 것을 확인할 수 있다.

```go
package main

func main() {
	test()
}

func test() {
	sliceA := make([]byte, 100)
	forReturn(sliceA)
}

func forReturn([]byte) {
	return
}
```

```
./main.go:12:6: can inline forReturn
./main.go:7:6: can inline test
./main.go:9:11: inlining call to forReturn
./main.go:3:6: can inline main
./main.go:4:6: inlining call to test
./main.go:4:6: inlining call to forReturn
./main.go:4:6: make([]byte, 100) does not escape
./main.go:8:16: make([]byte, 100) does not escape
```

sliceA는 make로 선언된 참조형이지만, 메모리 크기가 정해졌고 수명이 정해져있기 때문에 스택에 할당됐음을 알 수 있다.


```go
package main

func main() {
	_ = test()
}

func test() []byte {
	sliceA := make([]byte, 100)
	forReturn(sliceA)

	sliceB := make([]byte, 100)
	forReturn(sliceB)

	return sliceB
}

func forReturn([]byte) {
	return
}

```

```
./main.go:17:6: can inline forReturn
./main.go:7:6: can inline test
./main.go:9:11: inlining call to forReturn
./main.go:12:11: inlining call to forReturn
./main.go:3:6: can inline main
./main.go:4:10: inlining call to test
./main.go:4:10: inlining call to forReturn
./main.go:4:10: inlining call to forReturn
./main.go:4:10: make([]byte, 100) does not escape
./main.go:4:10: make([]byte, 100) does not escape
./main.go:8:16: make([]byte, 100) does not escape
./main.go:11:16: make([]byte, 100) escapes to heap
```

- sliceA는 forReturn() 으로 전달되지만, 수명이 정혀재있고 하위함수로 전달되기 때문에 스택에 할당된다.
- sliceB는 참조형으로, 상위함수에 전달되기 때문에 힙에 할당된다.

append 메모리 할당은 어떻게? / 익명함수일때는? (http://golang.site/go/article/11-Go-%ED%81%B4%EB%A1%9C%EC%A0%80) / 고루틴일때는? / 메서드일 때는? / 인터페이스일 떄는? / channel 공부 / map, slice

내가 짠 코드 or 쓰고 있는 아키텍처에는 어떻게 하는지? - 클린 아키텍처에서 메모리 관리를 어떻게 하는지?

그렇다면 메모리 관리를 잘 하기 위해서는 어떻게 해야하는가?

스택 힙 장단점 / 이스케이프 분석 특징 설정