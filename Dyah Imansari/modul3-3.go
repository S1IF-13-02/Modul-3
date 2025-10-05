package main

import "fmt"

func main() {
	var (
		Tahun   int
		Kabisat bool
	)
	fmt.Scan(&Tahun)
	fmt.Println("Tahun:", Tahun)
	Kabisat = (Tahun%400 == 0) || (Tahun%4 == 0) && (Tahun%100 != 0)
	fmt.Println("Kabisat:", Kabisat)

}
