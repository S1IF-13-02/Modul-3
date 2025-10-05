package main

import "fmt"

func main() {

	var year int

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0){
		fmt.Println(year, "adalah tahun kabisat")
	}else{
		fmt.Println(year, "bukan tahun kabisat")
	}
}