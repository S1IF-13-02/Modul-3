package main

import "fmt"

func IsKabisat(tahun int) bool {
	if tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0 {
		return true
	}

	return false
}

func main() {
	var tahun int

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	result := IsKabisat(tahun)

	fmt.Printf("Kabisat: %t\n", result)
}
