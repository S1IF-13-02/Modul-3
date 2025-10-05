package main

import (
	"fmt"
	"math"
)

func main() {

	var nilaiX float64

	fmt.Printf("Masukkan nilai x: ")
	_ , err := fmt.Scanln(&nilaiX)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	hasilfloorfloat := math.Floor(nilaiX)
	hasilfloorint := int(hasilfloorfloat)
	fmt.Printf("Hasil akhir: %d\n", hasilfloorint)
}