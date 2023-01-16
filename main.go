package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Check Input
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Not enough arguments")
		return
	}

	//Handle \n
	replace := strings.ReplaceAll(args[0], "\\n", "\n")
	//Split Input
	parts := strings.Split(replace, "\n")

	//Read .txt File
	bytes, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Println("No such file")
		return
	}

	//Split .txt File
	letters := strings.Split(string(bytes), "\n")

	for _, line := range parts { //this is used to handle newlines
		for i := 1; i < 9; i++ {
			for _, char := range line {
				os.Stdout.WriteString(letters[(char-32)*9+rune(i)])
				// fmt.Println(letters[(char-32)*9+rune(i)])
			}
			if line == "" {
				fmt.Println()
				break
			}
			os.Stdout.WriteString("\n")
		}
	}
}
