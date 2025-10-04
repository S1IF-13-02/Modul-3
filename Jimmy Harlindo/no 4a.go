package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	// Formula Konversi: F = C * (9/5) + 32
	// Menggunakan 9.0/5.0 untuk memastikan perhitungan floating point
	fahrenheit := (celsius * 9.0 / 5.0) + 32

	// Mencetak hasil sebagai bilangan bulat (sesuai contoh soal)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
}