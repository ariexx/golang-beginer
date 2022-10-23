package main

import "fmt"

// membuat func dengan interface kosong
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 0
	} else {
		return "ups"
	}
}

func main() {
	//set data dengan interface kosong
	var data = Ups(1)
	fmt.Println(data)
}
