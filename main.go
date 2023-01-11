package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
func main() {
	//Get user input
	args := os.Args[1:]
	if len(userInput) != 1 {
		fmt.Println("Not enough arguments")
		return
	}
	//Split input. Make the program accept \n as \\n
	parts := strings.Split(strings.ReplaceAll(args[0], "\\n", "\n"), "\n")

	//Read .txt fie
	bytes, err := ioutil.ReadFile("banner/standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Split .txt file
	letters := strings.Split(string(bytes), "\n")

	for _, line := range parts {
		for i := 1; i < 9; i++ {
			for _, char := range line {
				os.Stdout.WriteString(letters[(char-32)*9+rune(i)])
			}
			if line == "" {
				fmt.Println()
				break
			}
			os.Stdout.WriteString("\n")
		}
	}
}
*/

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	const fileName = "banner/standard.txt"

	fileContents, err := os.ReadFile(fileName) //mby take the args to shadow/standard/thinkertoy
	if err != nil {
		log.Fatalln(err)
	}

	fileLines := lineBreakRegExp.Split(string(fileContents), 2)

	for i, line := range fileLines {
		fmt.Printf("%d) \"%s\"\n", i+1, line)
	}
	//fmt.Println(fileLines)
}

//readfile
//map
/*
	maybe also use make? make is useful for looping
	drawings := map[srting]int{
		"lines 1 - 8":	1
		"lines 10-17": 	2
	}
*/

//For original one, use Map as a function
//in func main, scan user input, like at the very first fmt.scan.
//the values must match
//ReplaceString //n with /n

//to be added: args 3, choose from shadow/standard/thinkertoy(.txt)
