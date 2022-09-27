package main

import "fmt"

func sayHelloTo(fullName string) {
	fmt.Println("Halo", fullName)
}

func main() {
	sayHelloTo("Arief")
}