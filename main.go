package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//Get user input
	userInput := os.Args[1:]
	if len(userInput) != 1 {
		fmt.Println("Not enough arguments")
		return
	}
	//Split input. Make the program accept \n as \\n
	parts := strings.Split(strings.ReplaceAll(userInput[0], "\\n", "\n"), "\n")

	//Read .txt fie
	bytes, err := ioutil.ReadFile("banner/standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Split .txt file
	letters := strings.Split(string(bytes), "\n")
}

//to be added: args 3, choose from shadow/standard/thinkertoy(.txt)
