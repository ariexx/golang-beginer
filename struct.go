package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	arief := Customer {
		Name: "Arief",
		Address: "Indonesia",
		Age: 30,
	}

	anyone := Customer {
		Name: "Random",
		Address: "Aceh",
		Age: 30,
	}

	fmt.Println(arief, anyone)
}