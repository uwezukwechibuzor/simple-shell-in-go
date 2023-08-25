package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ") // print shell prompt

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSuffix(input, "\n")
		args := strings.Split(input, " ")

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Expected path")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(dir)
		case "ls":
			cmd := exec.Command("ls")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "date":
			fmt.Println(time.Now())
		case "whoami":
			cmd := exec.Command("whoami")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "mkdir":
			if len(args) < 2 {
				fmt.Println("Expected directory name")
				continue
			}
			err := os.Mkdir(args[1], 0755)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "rmdir":
			if len(args) < 2 {
				fmt.Println("Expected directory name")
				continue
			}
			err := os.Remove(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "rm":
			if len(args) < 2 {
				fmt.Println("Expected file name")
				continue
			}
			err := os.Remove(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "touch":
			if len(args) < 2 {
				fmt.Println("Expected file name")
				continue
			}
			_, err := os.Create(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "cp":
			if len(args) < 3 {
				fmt.Println("Expected source and destination")
				continue
			}
			cmd := exec.Command("cp", args[1], args[2])
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		case "mv":
			if len(args) < 3 {
				fmt.Println("Expected source and destination")
				continue
			}
			err := os.Rename(args[1], args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "hostname":
			host, err := os.Hostname()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Println(host)

		case "kill":
			if len(args) < 3 {
				fmt.Println("Expected signal and process ID")
				continue
			}
			signal, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			pid, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			err = process.Signal(syscall.Signal(signal))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		case "sleep":
			if len(args) < 2 {
				fmt.Println("Expected number of seconds")
				continue
			}
			duration, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			time.Sleep(time.Duration(duration) * time.Second)
		case "help":
			fmt.Println("Built-in commands: cd, pwd, ls, echo, date, whoami, mkdir, rmdir, rm, touch, cp, mv, hostname, kill, sleep, help, exit")
		case "exit":
			return
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
