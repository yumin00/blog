package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 파일을 binary로 만드는 함수
// 파일에 쓰는 프로세스
// 파일을 읽는 프로세스

const fifoName = "./myfifo"

func main() {
	makeFile()
	cmd := exec.Command("./main")
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing myfifo: %v\n", err)
		return
	}
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
