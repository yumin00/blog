package main

import (
	"fmt"
	"os"
)

func main() {
	const fifoName = "./myfifo"

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
