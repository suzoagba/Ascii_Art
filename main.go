package main

import (
	"fmt"
	"os"
)

func main() {
	userInput := os.Args[1:]
	if len(userInput) != 1 {
		fmt.Println("Not enough arguments")
		return
	}

	fmt.Println(userInput)
}
