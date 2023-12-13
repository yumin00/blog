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

#### [struct & 포인터]
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
./main.go:12:6: can inline test
./main.go:8:6: can inline main
./main.go:9:10: inlining call to test
./main.go:13:2: moved to heap: u
```

- `moved to heap: u`: test()에서 u는 user로 스택에 할당됐지만, u가 포인터형으로 반환되기 때문에 힙으로 이동된 것을 알 수 있다. Go 컴파일러는 기본적으로 가능한 한 변수를 스택에 할당하려고 시도하기 때문에 처음에 스택에 할당되는 것을 알 수 있다.

#### [moved to heap]

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


#### [slice]
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
- sliceB는 make로 선언된 참조형으로 상위함수에 전달되기 때문에 힙에 할당된다.

```go
package main

func main() {
	_ = test()
}

func test() []int {
	a := []int{1, 2, 3}
	return a
}
```
```
./main.go:7:6: can inline test
./main.go:3:6: can inline main
./main.go:4:10: inlining call to test
./main.go:4:10: []int{...} does not escape
./main.go:8:12: []int{...} escapes to heap
```

- array는 기본적으로 메모리 크기가 정해져있지 않기 때문에, 상위함수로 전달될 경우 힙에 할당된다.



#### [map]

```go
package main

func main() {
	test()
}

func test() map[string]string {
	user := map[string]string{
		"name": "yumin",
	}
	return user
}
```
```
./main.go:7:6: can inline test
./main.go:3:6: can inline main
./main.go:4:6: inlining call to test
./main.go:4:6: map[string]string{...} does not escape
./main.go:8:27: map[string]string{...} escapes to heap
```

- map(참조 타입)으로 할당할 경우, 크기가 정해지지 않기 때문에 map으로 할당된 변수가 return될 경우 힙에 할당된다.

```go
package main

func main() {
	test()
}

func test() string {
	user := map[string]string{
		"name": "yumin",
	}
	return user["name"]
}

```

````
./main.go:7:6: can inline test
./main.go:3:6: can inline main
./main.go:4:6: inlining call to test
./main.go:4:6: map[string]string{...} does not escape
./main.go:8:27: map[string]string{...} does not escape
````

- 하지만 map으로 할당될지라도(동적으로 할당되어 크기가 정해지지 않았더라도), 해당 변수가 직접 return 되지 않고 정적 변수가 return된다면 힙에 할당되지 않는다.

#### [slice]


[channel]
```go
package main

func main() {
	resultChan := makeChannel()

	go calculateSum(5, 10, resultChan)

	result := <-resultChan
	forReturn(resultChan)
	forReturnInt(result)

}

func makeChannel() chan int {
	resultChan := make(chan int)
	return resultChan
}

func calculateSum(a, b int, resultChan chan<- int) {
	sum := a + b
	resultChan <- sum
}

func forReturn(resultChan <-chan int) <-chan int {
	return resultChan
}

func forReturnInt(int) {
	return
}

```

```
./main.go:14:6: can inline makeChannel
./main.go:19:6: can inline calculateSum
./main.go:24:6: can inline forReturn
./main.go:28:6: can inline forReturnInt
./main.go:4:27: inlining call to makeChannel
./main.go:9:11: inlining call to forReturn
./main.go:10:14: inlining call to forReturnInt
./main.go:19:29: resultChan does not escape
./main.go:24:16: leaking param: resultChan to result ~r0 level=0
```

```go
package main

import (
	"time"
)

func main() {
	dataChan := make(chan int)

	go sendData(dataChan)

	for i := 1; i <= 5; i++ {
		data := <-dataChan
		forReturn(data)
	}

	close(dataChan)
}

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}

	close(ch)
}

func forReturn(int) {
	return
}
```

```
./main.go:20:6: can inline sendData
./main.go:29:6: can inline forReturn
./main.go:14:12: inlining call to forReturn
./main.go:20:15: ch does not escape
```

```go
package main

import (
	"time"
)

var DataChan chan int

func main() {
	go sendData(DataChan)

	for i := 1; i <= 5; i++ {
		data := <-DataChan
		forReturn(data)
	}

	close(DataChan)
}

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}

	close(ch)
}

func forReturn(int) {
	return
}
```

```
./main.go:20:6: can inline sendData
./main.go:29:6: can inline forReturn
./main.go:14:12: inlining call to forReturn
./main.go:20:15: ch does not escape
```

- channel을 사용했을 경우, 모두 스택에 할당됨을 확인할 수 있다.
- channel에 값을 넣어놓고 여러 개의 고루틴이 함께 사용하는 객체인데, 왜 힙이 아니라 스택에 할당될까?

채널이 힙이 아닌 스택에 할당되는 이유에 대해 공부해보기 위해 채널에 대해 자세히 공부해보고자 한다.

# Channel in GoLang
## Channel의 특징
### send / receive
고에서 채널은 여러 개의 고루틴이 동시에 접근하여 여러 데이터를 send할 때 누락되는 데이터가 존재하지 않는다.
또한, 여러 개의 고루틴이 동시에 접근하여 데이터를 receive할 때 중복으로 receive하는 데이터가 존재하지 않는다.

### block, unblock
고루틴이 block되기도 unblock되기도 한다.

- 채널 버퍼가 꽉 찬 상태에서 채널에 데이터를 send할 때 block된다.
- 채널 버퍼가 꽉 찬 상태에서 다른 고루틴이 데이터를 receive하면 데이터를 send하고자 하는 고루틴이 unblock되면서 데이터를 send할 수 있다.

## Channel의 생성
채널을 생성하고 메모리 할당을 보았을 때, 스택에 할당된다고 했다. 이에 대해 자세히 알기 위해서는 실제로 채널의 구조를 파악해야하 한다.

### hchan struct
```go
type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // size of the circular queue
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    elemsize uint16
    closed   uint32
    elemtype *_type // element type
    sendx    uint   // send index
    recvx    uint   // receive index
    recvq    waitq  // list of recv waiters
    sendq    waitq  // list of send waiters
    
    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
    lock mutex
}
```

```go
ch := make(chan int, 3)
```

위와 같이 채널을 생성하면 내부적으로 hchan struct가 만들어진다. hchan 고랭 내부에 정의된 것이다.


훅에서 메시지를 받아서 채널에 던져줌 . 채널에 넣어주고. 값을 꺼내서 핸들러 호출

- `qcount`: 큐에 저장된 데이터 총량을 의미한다.(ex. 0, 1, 2, 3)
- `dataqsiz`: 순환 큐의 크기를 나타낸다. 즉, 채널이 동시에 보유할 수 있는 요소의 최대 개수를 의미힌다. (ex. 3)
- `buf`: dataqsiz의 요소들의 배열의 포인터이다. 즉, 요소들을 저장하는 순환큐 배열의 포인터이다.
- `elemsize`: 큐에 저장되는 각 요소의 크기(바이트 단위)를 의미한다.
- `closed`: 채널이 닫혔는지에 대한 플래그이다. 1이면 채널이 닫혔다는 것을 의미한다.
- `elemtype`: 큐에 저장되는 요소의 타입을 나타내는 _type의 포인터이다.
- `sendx` / `recvx`: 채널에서 데이터를 받을 / 줄 배열의 인덱스이다.
- `recq` / `sendq`: 데이터를 받기를 / 쓰기를 기다리는 고루틴의 리스트이다. waitq 구조체를 사용한다.
- `lock`: hchan의 모든 필드를 보호하는 역할이다. lock일 경우 다른 고루틴의 상태를 변경하지 않아야 한다.


### 메모리 할당
채널은 스택에 할당된다. 그리고 채널이 생성된 후 그 내부의 hstruct는 힙에 할당된다. 채널은 hstruct의 포인터이다.
따라서 서로 다른 고루틴에서 쓰고 읽기가 가능한 것이다.

여러 개의 고루틴에서 각자의 스택에 채널을 할당해놓고, 실제로 데이터를 읽고 쓸 때는 값복사를 통해 hstruct가 있는 힙에 접근하여 읽고 쓰는 것이다.

## channel 코드 분석