package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	// cek kabisat
	isKabisat := false
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		isKabisat = true
	}

	// output
	fmt.Printf("Kabisat: %t\n", isKabisat)
}
