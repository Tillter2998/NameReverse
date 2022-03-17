package main

import (
	"NameReverse/helpers"
	"fmt"
	"log"
)

// test
func main() {
	reversedName, err := helpers.ReverseName("Elrich")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversedName)
}
