package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Input error:", err)
	}
	return in.Text()
}

func KillProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Kill()
}

func main() {
	fmt.Fprint(os.Stdout, "Shell v 1.0 ")
	sys, _ := os.LookupEnv("OS")
	fmt.Fprintln(os.Stdout, sys)
	var command string
	for command != "quit" {
		fmt.Fprint(os.Stdout, "shell->")
		command = Scan()
		arg := strings.Split(command, " ")
		cmd := strings.ToLower(arg[0])
		switch cmd {
		case "cd":
			//dir := "cd "+ arg[1]
			//exec.Command("cd", arg[1]).Run()
			os.Chdir(arg[1])
		case "pwd":
			/*			if sys == "Windows_NT" {
							cmd := exec.Command("cd")
							cmd.Stdout = os.Stdout
							cmd.Run()
						} else {
							cmd := exec.Command("pwd")
							cmd.Stdout = os.Stdout
							cmd.Run()
						}*/
			cur, _ := os.Getwd()
			fmt.Fprintln(os.Stdout, cur)
		case "echo":
			arg = arg[1:]
			cmd := exec.Command("echo", arg...)
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "kill":
			pid, _ := strconv.Atoi(arg[1])
			fmt.Fprintln(os.Stdout, KillProcess(pid))
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}

}
