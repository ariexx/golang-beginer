package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu di blok")
	}else{
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("danu", blacklist)

	/*jika namanya admin maka di blok*/
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}