package main

import "fmt"

func main () {
	var tahun int
	var kabisat bool

	fmt.Print("Masukan tahun :")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 {
		kabisat = true 
	} else if tahun%4 == 0  && tahun%100 !=0 {
		kabisat = true
	} else {
		kabisat = false 
	}

	fmt.Println("Tahun :",tahun)
	fmt.Println("kabisat:",kabisat)

}