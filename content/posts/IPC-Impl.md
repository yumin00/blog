---
title: "IPC Impl"
date: 2023-08-24T11:58:33+09:00
draft: true
---

IPC 구현 방법들을 직접 go 언어로 구현해보고자 한다.

## 1. PIPE
먼저, 파이프의 특징에 대해서 다시 정래해보자.
- 단방향 통신만을 지원
- 읽기를 위한 끝과 쓰기를 위한 끝이 있음
- 한 프로세스는 파이프의 끝에 데이터를 쓰고, 다른 한 프로세스는 파이프의 끝에서 데이터를 읽는다.
- 파이프는 내부적으로 버퍼를 사용하여 데이터를 저장한다. 버퍼가 가득 차면, 쓰기 작업은 블록된다.
- 프로세스가 종료되면 해당 데이터를 소멸된다.

````go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Println(w, "Hello world")
		w.Close()
	}()

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
````

파이프 특징을 바탕으로 golang으로 pipe를 구현해보았다.

1. `io.Pipe()`를 통해 파이프 하나를 구현하여 두 개의 파이프 끝, PipeReader와 PipeWriter를 반환한다.
2. 고루틴을 통해서 파이프의 쓰기 끝, w에 "Hello world"를 쓴다.
3. 메인 스레드에서는 `io.Copy(os.Stdout, r)`를 호출하여 파이프 읽기 끝, r에서 데이터를 읽어들여 `os.Stdout`으로 전송한다.
4. 파이프의 쓰기 끝, w가 닫히  `io.Copy()`는 반환하고 끝난다.

## 2. FIFO(named pipe)
