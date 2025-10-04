package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Daftar tahun dalam bentuk string (sesuai contoh soal)
	tahunStr := []string{"2016", "2000", "2018"}

	for _, bilangan := range tahunStr {
		// Konversi string -> int
		tahun, err := strconv.Atoi(bilangan)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Logika tahun kabisat
		isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

		// Konversi hasil ke string
		hasil := strconv.FormatBool(isKabisat)

		// Cetak output
		fmt.Println("Tahun: " + bilangan)
		fmt.Println("Kabisat: " + hasil)
		fmt.Println()
	}
}
