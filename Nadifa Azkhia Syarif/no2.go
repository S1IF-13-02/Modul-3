package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Input jejari dalam bentuk string
	var bilangan string = "5"

	// Konversi string ke integer
	teksBilangan, err := strconv.Atoi(bilangan)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// Hitung volume
		volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(teksBilangan), 3)

		// Hitung luas
		luas := 4 * math.Pi * math.Pow(float64(teksBilangan), 2)

		// Konversi hasil ke string
		hasilVolume := strconv.FormatFloat(volume, 'f', 4, 64)
		hasilLuas := strconv.FormatFloat(luas, 'f', 4, 64)

		// Cetak output
		fmt.Println("Bola dengan jejari " + bilangan +
			" memiliki volume " + hasilVolume +
			" dan luas kulit " + hasilLuas)
	}
}
