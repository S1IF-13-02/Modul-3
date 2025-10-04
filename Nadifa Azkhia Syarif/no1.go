package main

import (
	"fmt"
)

// Fungsi untuk menghitung x dari f(x)
func hitungX(fx float64) float64 {
	return (2 / (fx - 5)) - 5
}

func main() {
	// Daftar input sesuai contoh
	masukan := []float64{5.2, 4.125}

	fmt.Printf("%-5s %-10s %-10s\n", "No", "Masukan", "Keluaran")

	for i, fx := range masukan {
		x := hitungX(fx)
		fmt.Printf("%-5d %-10.3f %-10.3f\n", i+1, fx, x)
	}
}
