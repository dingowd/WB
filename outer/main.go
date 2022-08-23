package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	/*	for {
		fmt.Fprintln(os.Stdout, "I got control of the os.Stdout!!! Nobody can't stop me!!!")
	}*/
	env := os.Environ()
	binary, _ := exec.LookPath("outer.exe")
	args := []string{}
	err := syscall.Exec(binary, args, env)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
}
