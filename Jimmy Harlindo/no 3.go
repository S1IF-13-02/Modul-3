package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	// Logika Tahun Kabisat: (Habis dibagi 400) ATAU (Habis dibagi 4 TAPI TIDAK habis dibagi 100)
	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Printf("Kabisat: %t\n", isKabisat)
}