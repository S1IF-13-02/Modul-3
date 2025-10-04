package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	// Cek tahun kabisat
	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	// Tampilkan hasil
	fmt.Println("Tahun:", tahun)
	fmt.Println("Kabisat:", isKabisat)
}
