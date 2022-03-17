package main

import (
	"NameReverse/helpers"
	"fmt"
	"log"
)

func main() {
	reversedName, err := helpers.NameReverser("Elrich")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversedName)
}
