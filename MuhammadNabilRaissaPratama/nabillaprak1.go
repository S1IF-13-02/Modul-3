package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Masukkan x: ")
	fmt.Scan(&input)

	x, _ := strconv.Atoi(input)
	fx := (2 / (x - 5)) + 5
	fmt.Println("f(x) =", fx)
}
