package main

import (
	"fmt"
	"os/exec"
)

func main() {
	prg := "echo"

	arg1 := "there"
	arg2 := "are three"
	arg3 := "falcons"

	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
