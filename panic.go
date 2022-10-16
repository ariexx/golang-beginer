package main

import "fmt"

func endapp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endapp()
	if(error) {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}