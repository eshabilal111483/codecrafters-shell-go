package main

import (
	"bufio"
	"fmt"
	"os"
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
			case "echo", "exit", "type":
				fmt.Println(command[5:] + " is a shell builtin")
			default:
				//path, err := exec.LookPath("fortune")
				//if err != nil {
				//	log.Fatal("installing fortune is in your future")
				//}
				for _, p := range strings.Split(path, string(os.PathListSeparator)) {
					if strings.HasSuffix(p, "/"+command[5:]) {
						info, err := os.Stat(p)
						if err != nil {
							break
						}
						if info.Mode()&0111 == 0 {
							break
						}
					}
				}
				fmt.Println(command[5:] + ": not found")
			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
