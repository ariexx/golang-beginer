package main

import "fmt"

func main() {

	/*
	* creating new map
	*/
	person := map[string]string{
		"name":    "Arief",
		"address": "Medan",
	}
	fmt.Println(person)
	fmt.Println("Name ", person["name"])
	fmt.Println("Address ", person["address"])

	/*
	* Creating new key on person
	*/
	person["isEmployee"] = "iya"
	fmt.Println("is employee", person["isEmployee"])

	/*
	* Change is employee value
	*/
	person["isEmployee"] = "tidak"
	fmt.Println("is employee ? ", person["isEmployee"])

	/*
	* Delete data map
	*/
	delete(person, "isEmployee")
	fmt.Println(person)
	fmt.Println("Apakah datanya ada ? ", person["isEmployee"])
}