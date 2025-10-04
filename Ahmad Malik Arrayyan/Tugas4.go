package main
import (
	"fmt"
)
func main() {
	// Deklarasi variabel untuk input temperatur Celsius
	var celsius float64
	// Membaca input temperatur Celsius dari user
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)
	// Konversi ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32
	// Konversi ke Reamur
	reamur := celsius * 4 / 5
	// Konversi ke Kelvin
	kelvin := celsius + 273
	// Menampilkan hasil konversi
	fmt.Printf("Derajat Reamur : %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit : %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin : %.2f\n", kelvin)
}