package main

import (
	"fmt"
	"math"
)
// Fungsi untuk menghitung nilai x dari persamaan f(x) = 2/(x+5) + 5
// diberikan nilai f(x) (fx)
func calculateX(fx float64) (float64, error) {
	// Persamaan: fx = 2/(x+5) + 5
	// Maka: fx - 5 = 2/(x+5)
	// (x+5) = 2 / (fx-5)
	// x = (2 / (fx-5)) - 5
	// Cek pembagi tidak boleh nol
	if fx == 5 {
		return 0, fmt.Errorf("nilai f(x) tidak boleh sama dengan 5 (pembagian oleh nol)")
	}
	x := (2 / (fx - 5)) - 5
	return x, nil
}
func main() {
	// Contoh input f(x)
	inputs := []float64{5.2, 4.125}
	for _, fx := range inputs {
		x, err := calculateX(fx)
		if err != nil {
			fmt.Printf("Input f(x) = %.3f, Error: %s\n", fx, err)
		} else {
			// Membulatkan hasil ke tiga tempat desimal
			x = math.Round(x*1000) / 1000
			fmt.Printf("Input f(x) = %.3f, Output x = %.0f\n", fx, x)
		}
	}
}