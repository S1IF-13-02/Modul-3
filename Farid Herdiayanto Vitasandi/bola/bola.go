package main

import "fmt"

func main() {

	var r float64
	const phi = 3.14

	fmt.Print("Masukkan panjang jari-jari bola: ")
	fmt.Scan(&r)

	volume := (4.0/3.0) * phi * r * r * r
	luas := 4 * phi * r * r

	fmt.Println("Jari-jari bola adalah: ", r)
	fmt.Printf("Bola dengan jari-jari %.2f memiliki volume %.2f dan luas permukaan %.2f", r, volume, luas)
}
