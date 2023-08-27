package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
)

// 파일을 binary로 만드는 함수
// 파일에 쓰는 프로세스
// 파일을 읽는 프로세스

const fifoName = "./myfifo"

func main() {
	makeFile()
	writeFileProcess()
	readFileProcess()
}

func makeFile() {
	os.Remove(fifoName)
	err := syscall.Mkfifo(fifoName, 0666)
	if err != nil {
		log.Fatalf("Failed to create named pipe: %v", err)
	} else {
		fmt.Println("Named pipe created!")
	}
}

func readFileProcess() {
	f, err := os.Open(fifoName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("Received from named pipe:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func writeFileProcess() {
	f, err := os.OpenFile(fifoName, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Hello from named pipe!")
	if err != nil {
		panic(err)
	}

	fmt.Println("Message written to named pipe!")
}
