package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Encoding
func encode(input string) []int {
	var numbers []int

	for _, character := range input {
		var number int = int(character + 0)
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	//Scan the user input
	var input string
	fmt.Scanln(&input)

	//Encode the input
	var numbers []int = encode(input)

	//-----------------------
	//Reads the file
	bytes, _ := ioutil.ReadFile("banner/standard.txt")

	lines := strings.Split(string(bytes), "\n")
	x := numbers
	beginning := (x-1)*9 + 2
	ending := (x-1)*9 + 2 + 7
	for i, line := range lines {
		if i == 13 { //need to make a algorithm
			fmt.Println(line)
		}
	}
	//Encode the bytes by 8 lines

	//value of fist 8 bytes is 1
	//next 2 is 0
	// then 8 again

	//Read file
	//assign them values.
	// 8 long, 2 between, 8 long, 2 between
	//each is 8 bytes long, same order as the ascii tabel
	//replace the values
	fmt.Println(numbers)
	//fmt.Println(string(bytes))

}

// Read user input +
//Encode +
//Assign values to .txt files
//to Ascii Art
