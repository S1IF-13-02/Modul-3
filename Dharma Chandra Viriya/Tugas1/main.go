package main

import "fmt"

func MenyatakanX(fx float64) float64 {
	return (2.0 / (fx - 5.0)) - 5.0
}

func main() {
	var fx float64

	fmt.Print("Maasukkan bilangan f(x): ")
	fmt.Scan(&fx)

	result := MenyatakanX(fx)

	fmt.Println(int(result))
}
