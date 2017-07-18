package main

import (
	"strings"
	"os/exec"
	"fmt"
)

func main() {
	guilder := Guilder{}

	guilder.exe("docker run --net=host hekonsek/streamshift-zookeeper:0.0.1-SNAPSHOT")
	guilder.exe("docker run --net=host hekonsek/streamshift-kafka:0.0.1-SNAPSHOT")
	guilder.exeBackground("go run gojitsu-endpoint-rest.go")

	fmt.Println("Testing application...")

	guilder.exe("go build")
	guilder.exe("docker build . -t gojitsu/endpoint-rest:0.0.0-SNAPSHOT")
}

// Guilder build library

type Guilder struct {
}

func (g Guilder) exe(Command string)  {
	commandParts := strings.Split(Command, " ")
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if(err != nil) {
		fmt.Println(err)
	}
	if(stdoutStderr != nil) {
		fmt.Println(string(stdoutStderr))
	}
}

func (g Guilder) exeBackground(Command string) {
	commandParts := strings.Split(Command, " ")
	exec.Command(commandParts[0], commandParts[1:]...).Start()
}