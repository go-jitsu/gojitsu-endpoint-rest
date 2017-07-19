package main

import (
	"strings"
	"os/exec"
	"fmt"
)

func main() {
	guilder := Guilder{}

	guilder.ExecJoined("docker run -d --net=host hekonsek/streamshift-zookeeper:0.0.1-SNAPSHOT")
	guilder.ExecJoined("docker run -d --net=host hekonsek/streamshift-kafka:0.0.1-SNAPSHOT")

	guilder.AssumeExecNotContains("go test", "FAIL", "Tests succeeded.")

	guilder.ExecJoined("go build")
	guilder.AssumeExecContains("docker build . -t gojitsu/endpoint-rest:0.0.0-SNAPSHOT", "Successfully built", "Docker image created.")
	guilder.ExecJoined("docker push gojitsu/endpoint-rest:0.0.0-SNAPSHOT")
}

// Guilder build library

type Guilder struct {
}

func (g Guilder) ExecJoined(command string) ([]string, error)  {
	commandParts := strings.Split(command, " ")
	process := exec.Command(commandParts[0], commandParts[1:]...)
	stdoutStderr, err := process.CombinedOutput()
	return strings.Split(string(stdoutStderr), "\n"), err
}

func (g Guilder) AssumeExecContains(command string, assumeContains string, successMessage string) {
	output, err := g.ExecJoined(command)
	if(err != nil) {
		panic(err)
	}

	joinedOutput := strings.Join(output, "\n")
	if(!strings.Contains(joinedOutput, assumeContains)) {
		panic("The following output doesn't contain string " + assumeContains + `:\n` + joinedOutput)
	}

	fmt.Println(successMessage)
}

func (g Guilder) AssumeExecNotContains(command string, assumeNotContains string, successMessage string) {
	output, err := g.ExecJoined(command)
	if(err != nil) {
		panic(err)
	}

	joinedOutput := strings.Join(output, "\n")
	if(strings.Contains(joinedOutput, assumeNotContains)) {
		panic("The following output contains string " + assumeNotContains + `:\n` + joinedOutput)
	}

	fmt.Println(successMessage)
}