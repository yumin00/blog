---
title: "대표적인 자료구조에 대해 알아보고, golang으로 구현해보자"
date: 2023-09-22T12:21:14+09:00
draft: true
categories :
- DataStructure
---

# 자료구조(Data Structrue)
## 1. 자료구조란?
자료구조란, 데이터 값의 모임이다. 컴퓨터의 메모리 자원은 한정적인 반면 처리해야 할 데이터는 무수히 많을 수 있기 때문에 메모리 자원을 효율적으로 사용하기 위해 필요한 것이 바로 자료 구조이다.

## 2. 자료구조의 종류
### [배열(Array)]
- 동일한 타입의 데이터들을 저장하며, 고정된 크기를 가지고 있다.
- 인덱싱이 되어 있어, 인덱스 번호를 통해 데이터에 접근할 수 있다.

go에서 배열은 다음과 같이 표현할 수 있다.
```go
//1
var number []int
number = make([]int, 3)
number[0] = 1
number[1] = 2
number[2] = 3

//2
number := [3]int{1, 2, 3}

//3
var number = [3]int{1, 2, 3}
```


### [연결 리스트(Linked List)]
- 각 데이터 시퀀스가 순서를 가지고 연결된 순차적 구조
- 동적인 데이터 추가/삭제에 유리하다.

![image](https://github.com/yumin00/blog/assets/130362583/180e0dfc-7941-47fd-9632-97ed1104546f)
- 연결 리스트 구성 요소
  - Node : 각 요소
  - Key : 각 Node는 Key를 가지고 있음
  - Next : 각 Node는 다음 Node를 가리키는 포인터인 Next가 포함
  - Head : 첫 번째 Node
  - Tail : 마지막 Node
- ex) 프로그램 간 전환

go에서 연결 리스트는 다음과 같이 구현할 수 있다.

먼저 LinkedList struct와 Node struct를 만들어주고, Node를 생성하고 삭제하고 해당 LinkedList를 출력하는 코드를 구현했다.
```go
package main

import (
	"fmt"
)

type LinkedList struct {
	Head  *Node
	Tail  *Node
	count int
}

type Node struct {
	key  interface{}
	next *Node
}

func main() {
	linkedList := New()
	linkedList.AddNode("1번 노드")
	linkedList.AddNode("2번 노드")
	linkedList.PrintLinkedList()
}

func New() *LinkedList {
	return &LinkedList{
		Head:  nil,
		Tail:  nil,
		count: 0,
	}
}

func (link *LinkedList) AddNode(key interface{}) {
	newNode := &Node{
		key: key,
	}
	if link.Head == nil {
		link.Head = newNode
		link.Tail = newNode
	} else {
		link.Tail.next = newNode
		link.Tail = newNode
	}
	link.count++
}

func (link *LinkedList) RemoveFirstNode() {
	if link.Head == nil {
		fmt.Print("삭제할 수 있는 노드가 없습니다.")
		return
	} else if link.Head == link.Tail {
		link.Head = nil
		link.Tail = nil
	} else {
		link.Tail = link.Head.next
	}
	link.count--
}

func (link *LinkedList) RemoveLastNode() {
	if link.Tail == nil {
		fmt.Print("삭제할 수 있는 노드가 없습니다.")
	} else if link.Head == link.Tail {
		link.Head = nil
		link.Tail = nil
	} else {
		for node := link.Head; node != nil; {
			if node.next == link.Tail {
				link.Tail = node
			}
			node = node.next
		}
	}
	link.count--
}

func (link *LinkedList) PrintLinkedList() {
	current := link.Head
	for current != nil {
		fmt.Print("[", current.key, "]", "->")
		current = current.next
	}
	fmt.Println("nil")
}

```

### [스택(Stack)]
- 순서가 보존되는 선형 데이터 구조
- LIFO(Last In First Out)

go로 스택을 구현해보았다. 먼저 interface유형으로 stack을 만들고, stack의 마지막에 데이터를 append하거나 stack의 마지막 데이터를 pop하는 코드를 구현했다.

```go
package main

import "fmt"

type Stack []interface{}

func main() {
	var s Stack
	s.Push("1번")
	s.Push("2번")
	s.Push("3번")
	s.Pop()
	s.Pop()
	fmt.Println(s)
}

func (s *Stack) Push(data interface{}) {
	*s = append(*s, data) //stack의 마지막에 data 값 push
}

func (s *Stack) Pop() {
	if len(*s) == 0 {
		fmt.Println("stack is empty")
	} else {
		topIdx := len(*s) - 1
		topData := (*s)[topIdx]
		*s = (*s)[:topIdx]
	}
}
```

### [큐(Queue)]
- FIFO(First In First Out)
- ex) 멀티 스레딩에서 스레드 관리 / 대기열 시스템

go로 큐를 구현해보았다. 스택과 같이 append할 수 있도록 만들었고, 마지막 값을 pop하는 스택과 다르게 첫 번째 값을 pop할 수 있도록 코드를 구현했다.

```go
package main

import "fmt"

type Queue []interface{}

func (q *Queue) Push(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func main() {
	var q Queue
	q.Push("1번")
	q.Push("2번")
	q.Push("3번")
	q.Pop()
	q.Pop()
	fmt.Println(q)
}

```

### [해시 테이블(Hash Table)]
- 해시 함수를 사용하여 변환한 값을 index삼아 key와 value를 저장하는 자료구조
- 데이터의 크기와 상관없이 삽입 및 검색에 효율적이다.
- 충돌이 자주 일어날 수 있어 해시 함수를 개선하거나 테이블의 구조를 개선하는 chaining, open addressing 등의 방법이 사용된다.
- ex) db 인덱스 구현 / 사용자 로그인 인증

### [그래프(graph)]
- 노드 사이에 edge가 있는 형태
- directed graph : 일방통행
- undirected graph : 양방향
- ex) 소셜 미디어 네트워크 / 겁색 엔진에 의해 웹 링크를 나타내는 데 사용 / GPS에서 위치와 경로를 나타내는 데 사용

### [트리(tree)]
- 그래프가 계층적 구조를 가진 형태
- 최상위 노드(루트)를 가지고 있음
- 상위 노드를 부모 노드, 하위 노드를 자식 노드라고 부름

### [힙(heap)]
- Binary Tree(이진트리)
- 