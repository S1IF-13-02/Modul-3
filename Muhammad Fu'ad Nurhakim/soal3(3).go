package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	kabisat := tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0

	fmt.Println("tahun kabisat: ", kabisat)

}
