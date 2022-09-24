package main

import "fmt"

func main() {
	var nilaiAkhir = 30
	var absensi = 80

	// var lulusUjian = nilaiAkhir > 80
	// var lulusAbsensi = absensi > 30

	// var lulus = lulusUjian && lulusAbsensi

	fmt.Println(nilaiAkhir > 80 && absensi >= 80)
}