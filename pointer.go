package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address) //kosong karna new
	address4.City = "Medan"
	fmt.Println(address4)
}
