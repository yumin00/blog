---
title: "BOJ 11725 Golang 문제 풀이"
date: 2023-09-14T16:59:50+09:00
draft: false
categories :
- Algorithm
---

## 문제
- [BOJ 11725 트리의 부모 찾기](https://www.acmicpc.net/problem/11725)

메모리 초과와 시간 초과 때문에 애를 좀 먹으면서 푼 문제이다. 지금까지 문제를 풀 때, 메모리 초과나 시간 초과에 걸려본 적이 없었는데 처움으로 이러한 문제를 직면하여
메모리와 시간복잡도를 생각하여 문제를 풀어야하는 것을 알게 되었다.

## 풀이
### 1차 시도
처음에는 메모리 초과는 전혀 신경쓰지 않은 채, 미리 node를 만들고 dfs를 통해 문제를 풀었다.

node를 입력 받으면서 2차원을 배열을 만들었는데, (N+1)*(N+1)의 크기의 node였기 때문에 N이 100,000일 경우, 메모리를 초과하게 된다.
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(r, &N)

	node := make([][]int, N+1)

	for i := range node {
		node[i] = make([]int, N+1)
	}

	for i := 0; i < N-1; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		node[x][y] = 1
		node[y][x] = 1
	}

	for i := 2; i < N+1; i++ {
		for j := 0; j < len(node[i]); j++ {
			if node[i][j] == 1 {
				fmt.Fprintln(w, j)
				break
			} else {
				continue
			}
		}
	}
	w.Flush()
}
```

### 2차 시도
두 번째 시도에서는 메모리 초과를 해결하기 위해, bfs를 활용했다. node를 2차원 배열로 입력받고, node를 돌면서 q에 부모 노드와 자식 노드의 배열을 apeend하여 자식 노드의 부모를 출력하는 
방식으로 풀었다.

하지만 해당 풀이 방법은 시간 초과로 실패했다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(r, &N)

	node := make([][2]int, N-1)
	visited := make([]bool, N-1)
	var q [][]int

	for i := 0; i < N-1; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		node[i][0] = x
		node[i][1] = y
		if x == 1 {
			visited[i] = true
			q = append(q, []int{1, y})
		} else if y == 1 {
			visited[i] = true
			q = append(q, []int{1, node[i][0]})
		} else {
			continue
		}
	}

	lenghtQ := len(q)

	for i := 0; ; i++ {
		if i >= lenghtQ {
			break
		}
		for j := 0; j < N-1; j++ {
			if node[j][0] == q[i][1] && visited[j] == false {
				visited[j] = true
				q = append(q, []int{node[j][0], node[j][1]})
			}
			if node[j][1] == q[i][1] && visited[j] == false {
				visited[j] = true
				q = append(q, []int{node[j][1], node[j][0]})
			}
		}
		lenghtQ = len(q)
	}

	var result []int
	for i := 2; i < N+1; i++ {
		for j := 0; j < lenghtQ; j++ {
			if q[j][1] == i {
				result = append(result, q[j][0])
			}
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Fprintln(w, result[i])
	}
	w.Flush()
}
```

### 3차 시도
2차 시도에서 result를 append하는 과정에서 시간 초과가 발생했을 것이라고 판단했는데, q를 구하는 이중포문에서 result도 함께 append할 수 있는 로직을 생각해보았지만 마땅치 않아 새로운 로직을 
생각하여 다시 구현했다.

먼저 tree를 입력 받으면서, tree라는 이차원 배열에 각 노드에 연결된 노드를 넣어주었다.

bfs로 진행하기 위해 q를 만들고, 각 노드의 parents를 구하기 위해 parents 배열을 만들어주었다.

parents에서 1은 부모 노드가 없기 때문에 자기자신 1을 넣었고, q에는 1부터 bfs를 진행하기 위해 먼저 1을 넣었다.

for문을 통해 q의 length가 0이 될 때까지 돌면서, 각 노드의 부모 노드를 찾아 parents에 넣었고, parents의 index 2부터 출력하여 완성했다!
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	} else {
		return false
	}
}
func (q *Queue) EnQueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	var tree [][]int
	tree = make([][]int, n+1)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscanln(r, &x, &y)
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}

	var parents []int
	parents = make([]int, n+1)
	parents[1] = 1

	q := Queue{}
	q.EnQueue(1)

	for q.IsEmpty() == false {
		start := q.DeQueue().(int)
		for _, v := range tree[start] {
			node := v
			if parents[node] == 0 {
				parents[node] = start
				q.EnQueue(node)
			}
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintln(w, parents[i])
	}
	w.Flush()
}
```