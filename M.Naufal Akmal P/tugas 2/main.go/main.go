package main

import "fmt"

func main() {
	const phi = 3.1415926535
	var jejari float64
	fmt.Print("Masukkan nilai jejari: ")
	fmt.Scan(&jejari)

	volumeBola := (4.0 / 3.0) * phi * jejari * jejari * jejari
	luasBola := 4 * phi * jejari * jejari

	fmt.Print("bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f", jejari, volumeBola, luasBola)
}