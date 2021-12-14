package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	args := []string{
		"-c",
		"touch test-file-$(date +%H-%M-%S)",
	}
	cmd := exec.Command("bash", args...)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("combined output error: ", err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
