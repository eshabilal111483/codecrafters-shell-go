package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
//var _ = fmt.Print

func main() {

	//var command string
	reader := bufio.NewReader(os.Stdin)
	path := os.Getenv("PATH")

	for {

		// TODO: Uncomment the code below to pass the first stage
		fmt.Print("$ ")

		//_, err := fmt.Scan(&command)
		//if err != nil {
		//	fmt.Println("Error reading input:", err)
		//	return
		//}

		//fmt.Printf("%v: command not found", command)

		// Wait for user input
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			break

		} else if strings.Split(command, " ")[0] == "echo" {
			args := strings.Split(command, " ")[1:]
			fmt.Println(strings.Join(args, " "))

			// else if strings.HasPrefix(command, "echo ") {
			//	fmt.Println(command[5:])

		} else if strings.HasPrefix(command, "type ") {

			switch command[5:] {

			case "echo", "exit", "type", "pwd":
				fmt.Println(command[5:] + " is a shell builtin")

			default:
				//path, err := exec.LookPath("fortune")
				//if err != nil {
				//	log.Fatal("installing fortune is in your future")
				//}

				found := false
				dirs := strings.Split(path, string(os.PathListSeparator))

				for _, dir := range dirs {

					dir = strings.TrimSpace(dir)
					if dir == "" {
						continue
					}

					fullPath := filepath.Join(dir, command[5:])

					info, err := os.Stat(fullPath)
					if err != nil {
						continue
					}
					if info.Mode()&0111 != 0 {
						found = true
						fmt.Println(command[5:] + " is " + fullPath)
						break
					}
				}

				if !found {
					fmt.Println(command[5:] + ": not found")
				}

			}

		} else if command == "pwd" {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting pwd:", err)
				return
			}
			fmt.Println(pwd)

		} else if command == "cd" {
			// if aboslute path
			err := os.Chdir(strings.Split(command, " ")[1])
			if err != nil {
				fmt.Println("cd: ", strings.Split(command, " ")[1], ": No such file or directory.")
				return
			}

		} else {

			inputParts := strings.Split(command, " ")
			cmdName := inputParts[0]
			cmdArgs := inputParts[1:]

			found := false
			dirs := strings.Split(path, string(os.PathListSeparator))

			for _, dir := range dirs {

				dir = strings.TrimSpace(dir)
				if dir == "" {
					continue
				}

				fullPath := filepath.Join(dir, cmdName)

				info, err := os.Stat(fullPath)
				if err != nil {
					continue
				}
				if info.Mode()&0111 != 0 {
					found = true

					cmd := exec.Command(fullPath, cmdArgs...)
					cmd.Stdout = os.Stdout // Redirect output to your terminal
					cmd.Stderr = os.Stderr // Redirect errors to your terminal
					//cmd.Stdin = os.Stdin   // Forward user keystrokes if interactive

					cmd.Args = append([]string{cmdName}, cmdArgs...)

					err := cmd.Run()
					if err != nil {
						// Optional: handle runtime execution failures
					}

					break
				}
			}

			if !found {
				fmt.Println(command + ": command not found")
			}

		}

	}
}
