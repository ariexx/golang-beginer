package main

import "fmt"

func getFullName() (string, string, string) {
	return "Arief", "Muhammad", "Ganteng"
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}