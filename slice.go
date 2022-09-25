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

	/*
	* Membuat copy slice
	*/
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copy slice", copySlice)

	/*
	* Creating array or slice
	*/
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println("ini adalah array", iniArray, "ini adalah slice", iniSlice)
}