package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHi(customer Customer, name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

// struct function
func (customer Customer) sayMe(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	arief := Customer{
		Name:    "Arief",
		Address: "Indonesia",
		Age:     30,
	}

	anyone := Customer{
		Name:    "Random",
		Address: "Aceh",
		Age:     30,
	}

	sayHi(anyone, "Arief")
	anyone.sayMe("?")

	fmt.Println(arief, anyone)
}
