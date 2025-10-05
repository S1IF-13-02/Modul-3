package main

import "fmt"

func main() {
	const Dollar = 15000
	var nilaiRupiah int
	fmt.Print("Masukkan nilai rupiah: ")
	fmt.Scan(&nilaiRupiah)

	convert := nilaiRupiah / Dollar
	fmt.Printf("Hasilnya adalah %d", convert)
}