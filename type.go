package main

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var name = Name{
	Firstname: "Nikhil",
	Lastname:  "Sharma",
	Age:       20,
}
