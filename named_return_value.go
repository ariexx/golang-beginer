package main

import "fmt"

func getFullName2() (firstName string, lastName string, optional string) {
	firstName = "Arief"
	lastName = "Muhammad"
	optional = "Ganteng"

	return
}

func main() {
	firstName, lastName, optional := getFullName2()

	fmt.Println(firstName, lastName, optional)
}