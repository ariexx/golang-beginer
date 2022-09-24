package main

import "fmt"

func main() {
	/*
	* Membuat Inisialisasi Array
	*/
	var names [4]string

	names[0] = "Muhammad"
	names[1] = "Arief"
	names[2] = "Ganteng"
	names[3] = "kanabawi"

	fmt.Println(names)

	/*
	* Membuat array secara langsung
	*/
	var values = [3] uint8 {
		90,
		30,
		50,
	}

	fmt.Println(values)

	fmt.Println(len(values))
}