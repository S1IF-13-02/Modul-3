package main

import (
	"fmt"
)

func main() {
	// Input suhu dalam Celsius
	var celsius float64
	fmt.Print("Masukkan temperatur Celsius: ")
	fmt.Scanln(&celsius)

	// Konversi ke Reamur, Fahrenheit, Kelvin
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	// Output hasil konversi
	fmt.Println("Temperatur Celsius :", celsius)
	fmt.Println("Derajat Reamur     :", reamur)
	fmt.Println("Derajat Fahrenheit :", fahrenheit)
	fmt.Println("Derajat Kelvin     :", kelvin)
}
