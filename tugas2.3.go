package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("masukan tahun: ")
	fmt.Scan(&tahun)

	var kabisat bool
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		kabisat = true
	} else {
		kabisat = false
	}
	fmt.Print("masukan kabisat: ", kabisat)

}
