package main

import "fmt"

/*
* Function sumAll berfungsi untuk menjumlahkan beberapa int
* varargs (variadic argument) hanya bisa di akhir
 */
func sumAll(numbers ...int) int {
	total := 0
	// _ = ignore
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	total := sumAll(10, 20, 20, 30, 40) //return 120
	fmt.Println(total)

	slice := []int{10,20,30,40,50} //slice untuk ditambahkan kedalam variadic function
	total = sumAll(slice...) //merubah slice menjadi variadic
	fmt.Println(total) //totalkan semuanya
}