package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	tahunKabisat := tahun%400 == 0 || tahun%4 == 0
	bukanKabisat := tahun%100 != 0

	var kabisat bool = tahunKabisat && bukanKabisat
	fmt.Printf("Kabisat: %t", kabisat)
}