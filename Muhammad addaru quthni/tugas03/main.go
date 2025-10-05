package main

import "fmt"

func main() {
	var year int

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("Tahun %d adalah tahun kabisat.(true)\n", year)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat.(false)\n", year)
	}
}
