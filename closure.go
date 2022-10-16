package main

import "fmt"

func main() {
	name := "arief"
	counter := 0

	increment := func() {
		name := "Budi"
		name = "Arief geming"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}