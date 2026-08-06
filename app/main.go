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
		}

		fmt.Println(command[:len(command)-1] + ": command not found")

	}
}
