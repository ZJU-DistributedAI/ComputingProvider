package main

import (
	"fmt"
	"os/exec"
)

func main() {

	command := `./_test.sh`
	fmt.Println("cmd")
	cmd := exec.Command("/bin/bash", command)
	fmt.Println("output")
	output, err := cmd.Output()
	fmt.Println("Printf")
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
