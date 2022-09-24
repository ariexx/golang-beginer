package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"December",
	}

	var slice1 = months[4:7]

	fmt.Println("Length", len(slice1))
	fmt.Println("Cap", cap(slice1))

	/*
	* length 2
	* capacity 5
	*/

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Arief"
	newSlice[1] = "Ganteng"

	fmt.Println(newSlice)
	fmt.Println("Length ", len(newSlice))
	fmt.Println("Capacity ",cap(newSlice))
}