package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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
	cur, _ := os.Getwd()
	os.Chdir(cur)
	for command != "quit" {
		fmt.Fprintln(os.Stdout)
		fmt.Fprint(os.Stdout, "shell->")
		command = Scan()
		arg := strings.Split(command, " ")
		cmd := strings.ToLower(arg[0])
		switch cmd {
		case "cd":
			if len(arg) < 2 {
				continue
			}
			os.Chdir(arg[1])
		case "pwd":
			cur, _ := os.Getwd()
			fmt.Fprintln(os.Stdout, cur)
		case "echo":
			arg = arg[1:]
			cmd := exec.Command("echo", arg...)
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "kill":
			pid, _ := strconv.Atoi(arg[1])
			if err := KillProcess(pid); err == nil {
				fmt.Fprintln(os.Stdout, "Process with PID", pid, "killed")
			} else {
				fmt.Fprintln(os.Stdout, "Unable to kil process with PID", pid, "Error:", err.Error())
			}
		case "ps":
			if strings.ToLower(sys) == "windows_nt" {
				cmd := exec.Command("tasklist")
				cmd.Stdout = os.Stdout
				cmd.Run()
			} else {
				cmd := exec.Command("ps", arg[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		case "exec":
			env := os.Environ()
			binary, err := exec.LookPath(arg[1])
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				continue
			}
			if len(arg) == 1 {
				continue
			}
			args := make([]string, 0)
			if len(arg) >= 3 {
				args = arg[2:]
			} else {
				args = append(args, "")
			}
			if err := syscall.Exec(binary, args, env); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			}
		case "fork":
			binary, err := exec.LookPath(arg[1])
			attr := syscall.ProcAttr{
				"",
				[]string{},
				[]uintptr{},
				nil,
			}
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				continue
			}
			if len(arg) == 1 {
				continue
			}
			args := make([]string, 0)
			if len(arg) >= 3 {
				args = arg[2:]
			} else {
				args = append(args, "")
			}
			if _, err := syscall.ForkExec(binary, args, &attr); err != nil { // nolint: gocritic
				fmt.Fprintln(os.Stdout, err.Error()) // nolint: gocritic
			}
		default:
			args := make([]string, 0)
			if len(arg) == 1 {
				args = append(args, "")
			} else {
				arg = arg[1:]
			}
			cmd := exec.Command(cmd, args...)
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}
