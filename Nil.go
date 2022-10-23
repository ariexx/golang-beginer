package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	//tanpa nil
	var orang map[string]string

	if orang["name"] == "" {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(orang)
	}

	//check error dengan nil
	person := NewMap("Arief")
	fmt.Println(person)
}
