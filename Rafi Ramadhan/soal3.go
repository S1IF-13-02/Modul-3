package main

import "fmt"

func main() {
	var year int
	fmt.Print("Tahun: ")
	fmt.Scan(&year)

	isLeap := (year%400 == 0) || (year%4 == 0 && year%100 != 0)

	fmt.Printf("Kabisat: %t\n", isLeap)
}
