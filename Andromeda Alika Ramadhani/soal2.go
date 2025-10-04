package main

import "fmt"

func main() {
	var jari2 int
	const pi = 3.14
	fmt.Print("Masukkan jari-jari= ")
	fmt.Scan(&jari2)

	volume := (4.0 / 3.0) * pi * float64(jari2*jari2*jari2)
	luas := 4 * pi * float64(jari2*jari2)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas %.4f", jari2, volume, luas)
}
