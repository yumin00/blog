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

### Stack
- 함수 내 선언된 지역 변수는 기본적으로 스택에 할당 - 함수 호출이 끝날 때 해당 변수의 수명은 끝남
- 변수가 함수 내에서만 사용되고, 함수 호출이 끝날 때 해당 변수의 수명이 끝난다면 해당 변수는 스택에 할당됨

### Heap
- 이스케이프 분석: 함수 외부로 탈출하는 변수를 자동으로 힙에 할당
- 포인터를 사용하여 변수를 참조할 경우 힙에 할당
- `new`, `make` 같은 함수를 통해 동적으로 메모리를 할당하는 경우, 힙에 할당
- 가비지 컬렉터가 힙 메모리를 자동으로 관리해주기 때문에 개발자가 수동으로 메모리를 해제할 필요가 없음

### Go 메모리 할당
- 힙 할당은 무겁다고 한다. 그 이유는 할당 후 가비지 컬렉터가 할당된 객체 참조 여부를 주기적으로 검사하기 때문이다.
- 스택 할다은 가볍다고 한다. 스택 할당의 경우 CPU 명령어 PUSH/POP 두 개로 끝나기 때문이다.

## 이스케이프 분석
Go 컴파일러는 이스케이프 분석이라는 기술을 사용하여 스택 할당과 힙 할당 중 어떤 것을 사용할지 선택한다.

이스케이프 분석이란, 객체의 포인터가 서브 루틴 밖으로 전파되는지를 분석하는 기술이다.

### TEST
이스케이프 분석을 통해 직접 Go에서 스택, 힙을 어떻게 할당하는지에 대해 테스트를 해보았다.

```go
package escape

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
}

```

해당 코드를 `go build -gcflags ‘-m’` 같이 옵션을 지정하여 build 하면 이스케이프 분석 결과를 확인할 수 있다. 이스케이프 분석 결과는 다음과 같다.

```
./main.go:7:13: inlining call to fmt.Println
./main.go:7:13: ... argument does not escape
./main.go:7:14: x escapes to heap
```
여기에서 escaping이란, 변수가 스택 메모리 영역을 벗어나 힙 메모리에 할당됨을 의미한다.

해당 코드에서 x는 main 함수에서 사용되고 있고, 전역변수가 아니기 때문에 스택에 할당될 것이라고 생각했다.

x는 fmt.Println 라는 함수의 인수로 전달되고, 그 인수가 탈출하기 때문에 x도 탈출한다는 것을 알 수 있다.

```go
package escape

import "fmt"

type user struct {
	name string
}

func main() {
	v := "hi yumin"
	var a = v
	test(a)
	testPointer()
	fmt.Print(v, a)
}

func test(input string) {
	var temp string
	temp = input
	fmt.Print(temp)
}

func testPointer() {
	temp := &user{"yumin"}
	fmt.Println(temp)
}
```

```shell
./main.go:20:11: inlining call to fmt.Print
./main.go:25:13: inlining call to fmt.Println
./main.go:14:11: inlining call to fmt.Print
./main.go:17:11: leaking param: input
./main.go:20:11: ... argument does not escape
./main.go:20:12: temp escapes to heap
./main.go:24:10: &user{...} escapes to heap
./main.go:25:13: ... argument does not escape
./main.go:14:11: ... argument does not escape
./main.go:14:12: v escapes to heap
./main.go:14:15: a escapes to heap
```

- leaking param: input: 함수 test의 매개변수 input이 함수 범위 밖으로 "누출"(escaping)한다는 것을 의미한다. 여기서는 temp에 할당되고 fmt.Print로 전달되기 때문이다.
- temp escapes to heap: temp 변수가 힙에 할당되는 것을 의미한다. 이는 fmt.Print 호출에서 변수가 사용되기 때문이다.
- &user{...} escapes to heap: user 구조체의 포인터가 힙에 할다된다. 이는 구조체 포인터가 함수 범위를 벗어나서 사용되기 때문이다.
- ... argument does not escape : 여기서 ...는 fmt.Print 및 fmt.Println 함수에 전달되는 인자들이 함수 범위 밖으로 누출되지 않음을 나타낸다.
- v/a escapes to heap: v와 a 변수가 힙에 할당된다. 이는 이 변수들이 fmt.Print 호출에 사용되기 때문이다.