package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter /* importing Filter to function */) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	if name == "babi" {
		return "..."
	}

	if name == "babi jelek" {
		return "..."
	}

	if name == "jancok" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFilter("Arief", spamFilter)
	sayHelloWithFilter("babi jelek", spamFilter)
}