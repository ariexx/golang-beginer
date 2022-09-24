package main

import "fmt"

func main() {
	type NoNik string
	type Married bool

	var NoNikUser NoNik = "123123123123"
	var isMarried Married = true
	
	fmt.Println(NoNikUser, isMarried)
}