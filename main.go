package main

import (
	"fmt"
	"log"

	pb "github.com/NikhilSharmaWe/test/proto"
)

var nameBytes = name.ComposeNameBinary()

var naam = pb.Name{
	Firstname: "Nikhil",
	Lastname:  "Sharma",
	Age:       20,
}

var naamBytes, _ = ComposeNamePB(naam)

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

	b, err = ComposeNamePB(naam)
	if err != nil {
		log.Fatal(err)
	}

	x, err := ParseNamePB(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
}
