---
title: "[BOJ] 2606 바이러스 Golang 문제 풀이"
date: 2023-09-08T19:36:58+09:00
draft: false
categories :
- Algorithm
---

## 문제
- BOJ 2606: 바이러스 https://www.acmicpc.net/problem/2606

## 풀이
- BFS(너비 우선 탐색) 을 통해 구현했다.

1. 컴퓨터 수(count)와 연결되어 있는 컴퓨터 쌍(n)을 입력 받는다.
2. computer 를 2차원 배열로 만들고, 연결되어 있는 컴퓨터 번호 쌍을 입력 받는다.
3. 먼저 1과 연결된 computer를 찾기 위해 computer 를 순회하면서, 1과 연결되어 있는 컴퓨터를 q에 넣는다.
4. q에 있는 컴퓨터와 연결된 컴퓨터를 찾기 위해, q를 순회하면서 computer에 해당 q값이 있고 / 해당 컴퓨터와 연결되어 있는 컴퓨터의 번호가 q에 없으면 해당 컴퓨터의 번호를 q에 넣는다.
5. q를 모두 순회하면 break 하여 q의 length를 출력한다.

q를 순회하면서 이미 확인했던 computer 쌍을 확인하는 것을 막기 위해,

check 라는 배열을 만들고, 순회한 computer 의 index 는 true로 만들어서 true일 경우 해당 index는 continue 할 수 있도록 했다.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue []interface{}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var count int
	var n int
	fmt.Fscanln(r, &count)
	fmt.Fscanln(r, &n)

	computer := make([][2]int32, n)
	check := make([]bool, n)

	for i := 0; i < n; i++ {
		var a, b int32
		fmt.Fscanln(r, &a, &b)
		computer[i][0] = a
		computer[i][1] = b
	}

	for i := 0; i < n; i++ {
		check[i] = false
	}

	var q Queue

	for i := 0; i < n; i++ {
		if computer[i][0] == 1 {
			q = append(q, computer[i][1])
			check[i] = true
		} else if computer[i][1] == 1 {
			q = append(q, computer[i][0])
			check[i] = true
		} else {
			check[i] = false
			continue
		}
	}

	for i := 0; ; i++ {
		lengthQ := len(q)
		if i >= lengthQ || lengthQ == 0 {
			break
		}

		for j := 0; j < n; j++ {
			if check[j] == true {
				continue
			}
			if computer[j][0] == q[i] {
				isDuplicate := false
				for v := 0; v < lengthQ; v++ {
					if computer[j][1] == q[v] {
						isDuplicate = true
						break
					}
				}
				if isDuplicate == false {
					q = append(q, computer[j][1])
					check[j] = true
				}
			} else if computer[j][1] == q[i] {
				isDuplicate := false
				for v := 0; v < lengthQ; v++ {
					if computer[j][0] == q[v] {
						isDuplicate = true
						break
					}
				}
				if isDuplicate == false {
					q = append(q, computer[j][0])
					check[j] = true
				}
			}
		}
	}
	fmt.Fprintln(w, len(q))
	w.Flush()
}

```
