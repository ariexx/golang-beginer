package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	var hello = getHello("Arief")
	fmt.Println(hello)

	helloKosong := getHello("")
	fmt.Println(helloKosong)
}