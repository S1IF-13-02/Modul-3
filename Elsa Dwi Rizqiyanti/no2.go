package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Print("Masukkan jejari bola: ")
	fmt.Scan(&r)

	// Hitung volume dan luas bola
	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(r), 3)
	luas := 4 * math.Pi * math.Pow(float64(r), 2)

	// Tampilkan hasil dengan 4 angka di belakang koma
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
