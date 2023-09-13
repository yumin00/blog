---
title: "[BOJ] 1260 DFS와 BFS Golang 문제 풀이
date: 2023-09-11T19:11:36+09:00
draft: false
categories :
- Algorithm
---

## 문제
- BOJ 1260: DFS와 BFS 

## 풀이
1. 정점의 개수(N), 간선의 개수(M), 탐색을 시작할 정점의 번호(V)를 입력 받는다.
2. 입력받을 두 정점의 번호를 넣을 graph를 만들어준다.
3. visit을 체크하기 위해 visited도 만들어준다.
4. dfs를 진행한다.
5. visited를 reset해준다.
6. bfs를 진행한다.
7. 결과값을 출력한다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n       int
	graph   [][]int
	visited []bool
	N, M, V int
	r       *bufio.Reader
	w       *bufio.Writer
)

func main() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &N, &M, &V)

	graph = make([][]int, N+1)
	visited = make([]bool, N+1)

	for i := range graph {
		graph[i] = make([]int, N+1)
	}

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		graph[x][y] = 1
		graph[y][x] = 1
	}

	dfs(V)
	w.Flush()
	resetVisited()
	fmt.Println()

	bfs(V)
}

func dfs(V int) {
	visited[V] = true
	fmt.Fprint(w, V, " ")
	for i := 0; i < len(graph[V]); i++ {
		if graph[V][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}

func bfs(V int) {
	visited[V] = true
	q := []int{V}

	for len(q) != 0 {
		front := q[0]
		fmt.Print(front, " ")
		q = q[1:]
		for i := 0; i <= N; i++ {
			if graph[front][i] == 1 && !visited[i] {
				visited[i] = true
				q = append(q, i)
			}
		}
	}
}

func resetVisited() {
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
}

```