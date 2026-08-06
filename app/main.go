package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
//var _ = fmt.Print

func main() {

	var command string

	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("%v: command not found", command)

}
