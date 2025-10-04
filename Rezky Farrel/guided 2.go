package main

import "fmt"

func main() {
	var alas, tinggi float64
	fmt.Print("Masukkan nilai alas dan nilai tinggi: ")
	fmt.Scan(&alas, &tinggi)

	luasSegitiga := 0.5 * alas * tinggi
	fmt.Print("Luas segitiganya adalah %.1f", luasSegitiga)
}