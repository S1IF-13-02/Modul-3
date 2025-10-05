package main

import "fmt"

func main() {
	const π = 3.14
	var r float64

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * π * r * r * r
	luas := 4 * π * r * r

	fmt.Println("Volume bola =", volume)
	fmt.Println("Luas permukaan bola =", luas)
}
