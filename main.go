package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 { // require 2 arguments
		if os.Args[1] == "" {
			return
		}
		fmt.Println("Invalid Input. Usage: go run main.go \"<text>\" <font>")
		return
	}

	text := os.Args[1] // Get Text from Input
	font := os.Args[2] // Get Font File from Input

	if text == "\\n" { // convert "\n" to newline
		fmt.Printf("\n")
	} else if text != "" {
		fontBytes, err := os.ReadFile(font) // Read Font File
		if err != nil {
			fmt.Print("Error reading font file: ")
			return
		}

		lines := strings.Split(string(fontBytes), "\n") // Split Font File
		text = strings.ReplaceAll(text, "\\n", "\n")    // Handle "\n" input
		parts := strings.Split(text, "\n")              // Split Input

		for _, line := range parts {
			for i := 1; i < 9; i++ {
				for _, ascii := range line {
					fmt.Printf(lines[(ascii-32)*9+rune(i)])
				}
				if line == "" { // handle empty line ("\n")
					fmt.Println()
					break
				}
				fmt.Println()
			}
		}
	}
}
