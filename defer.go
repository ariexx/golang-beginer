package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil logging")
}

func runApp(value int) {
	defer logging() //eksekusi walau aplikasi error
	result := 10 / value
	fmt.Println("Result", result)
	fmt.Println("Run application")
}

func main() {
	runApp(2)
}