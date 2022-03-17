package main

import (
	"NameReverse/helpers"
	"fmt"
	"log"
)

// test
func main() {
	fmt.Print("Enter string to be reversed: ")
	var input string
	fmt.Scanln(&input)

	reversedName, err := helpers.ReverseString(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversedName)
}
