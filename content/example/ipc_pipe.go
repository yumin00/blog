package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("명령어, 파이프, 명령어 순으로 입력하세요.")
	fmt.Println("ex) echo \"hello world\" | grep \"hello\"")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	cmd := strings.Split(input, "|")
	scanCmd1 := cmd[0]
	scanCmd2 := cmd[1]

	cmd1, arg1 := splitCmd(scanCmd1)
	cmd2, arg2 := splitCmd(scanCmd2)

	// This will simulate "echo 'Hello World' | grep Hello"
	execCmd1 := exec.Command(cmd1, arg1)
	execCmd2 := exec.Command(cmd2, arg2)

	var outputBuf bytes.Buffer
	//cmd1의 출력값, Hello Wolrd를 버퍼에 저장
	execCmd1.Stdout = &outputBuf

	if err := execCmd1.Start(); err != nil {
		fmt.Println("Error starting cmd1:", err)
		return
	}
	if err := execCmd1.Wait(); err != nil {
		fmt.Println("Error waiting for cmd1:", err)
		return
	}

	// cmd2 입력값은 Hello World(cmd1의 출력값)
	execCmd2.Stdin = &outputBuf

	var finalOutputBuf bytes.Buffer

	//cmd2의 출력값, Hello World를 버퍼에 저장
	execCmd2.Stdout = &finalOutputBuf

	if err := execCmd2.Start(); err != nil {
		fmt.Println("Error starting cmd2:", err)
		return
	}
	if err := execCmd2.Wait(); err != nil {
		fmt.Println("Error waiting for cmd2:", err)
		return
	}

	//finalOutputBuf, Hello Wolrd 복사
	io.Copy(os.Stdout, &finalOutputBuf)
}

func splitCmd(scanCmd string) (string, string) {
	parts := strings.Split(scanCmd, "\"")

	cmd := strings.TrimSpace(parts[0])
	arg := parts[1]

	return cmd, arg
}

func basic() {
	// This will simulate "echo 'Hello World' | grep Hello"
	cmd1 := exec.Command("echo", "Hello World")
	cmd2 := exec.Command("grep", "Hello")

	var outputBuf bytes.Buffer
	//cmd1의 출력값, Hello Wolrd를 버퍼에 저장
	cmd1.Stdout = &outputBuf

	if err := cmd1.Start(); err != nil {
		fmt.Println("Error starting cmd1:", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println("Error waiting for cmd1:", err)
		return
	}

	// cmd2 입력값은 Hello World(cmd1의 출력값)
	cmd2.Stdin = &outputBuf

	var finalOutputBuf bytes.Buffer

	//cmd2의 출력값, Hello World를 버퍼에 저장
	cmd2.Stdout = &finalOutputBuf

	if err := cmd2.Start(); err != nil {
		fmt.Println("Error starting cmd2:", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Println("Error waiting for cmd2:", err)
		return
	}

	//finalOutputBuf, Hello Wolrd 복사
	io.Copy(os.Stdout, &finalOutputBuf)
}
