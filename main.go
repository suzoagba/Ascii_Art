package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args

	const usage = "Usage: 'go run main.go [TEXT]' or 'go run main.go [TEXT] [FONT]' "
	if len(args) < 2 || len(args) > 3 {
		fmt.Println(usage)
		return
	}

	text := args[1]
	font := "fonts/standard.txt"

	if len(args) == 3 {
		switch args[2] {
		case "standard", "standard.txt":
			font = "fonts/standard.txt"
		case "shadow", "shadow.txt":
			font = "fonts/shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			font = "fonts/thinkertoy.txt"
		default:
			fmt.Println("Invalid font type")
			return
		}
	}

	if text == "\\n" {
		fmt.Printf("\n")
	} else if text != "" {
		//Read .txt File
		fontBytes, err := ioutil.ReadFile(font)
		if err != nil {
			fmt.Println("No such file")
			return
		}

		// Split .txt File
		letters := strings.Split(string(fontBytes), "\n")

		//Handle \n
		text = strings.ReplaceAll(text, "\\n", "\n")

		//Split Input
		parts := strings.Split(text, "\n")

		for _, line := range parts {
			for i := 1; i < 9; i++ {
				for _, char := range line {
					os.Stdout.WriteString(letters[(char-32)*9+rune(i)])
				}

				os.Stdout.WriteString("\n")
			}
		}
	}
}
