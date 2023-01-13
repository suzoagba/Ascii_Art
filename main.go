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
	letters := strings.Split(string(bytes), "\n") // \n is the separator

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
/*
var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	const fileName = "banner/standard.txt"

	fileContents, err := os.ReadFile(fileName) //mby take the args to shadow/standard/thinkertoy
	if err != nil {
		log.Fatalln(err)
	}

	fileLines := lineBreakRegExp.Split(string(fileContents), 27)

	for i, line := range fileLines {
		fmt.Printf("%d) \"%s\"\n", i+1, line)
		chunks := Split(line, 3) / 3
		fmt.Printf(chunks)
	}
	//fmt.Println(fileLines)
}

func Split() {
	panic("unimplemented")
}
*/
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
