package main

import (
	"fmt"
	"log"
)

var nameBytes = name.ComposeNameBinary()

func main() {
	b, err := name.ComposeNameJSON()
	if err != nil {
		log.Fatal(err)
	}

	n, err := ParseNameJSON(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

	b = name.ComposeNameBinary()

	n = ParseNameBinary(b)

	fmt.Println(n)
}
