package main

import "fmt"

func endApp() {
	message := recover() //simpan fungsi recover didalam defer function
	if message != nil {
		fmt.Println("error", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}