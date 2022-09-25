package main

import "fmt"

func main() {

	name := "Arief"

	if name == "Arief" {
		fmt.Println("Hello Arief")
	} else if name == "eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, Boleh kenalan ?")
	}


	/*
	* Short statement for if
	*/
	
	if length := len(name) ;length > 5 {
		fmt.Println("Terlalu Panjang")
	}else{
		fmt.Println("Nama sudah benar")
	}
}