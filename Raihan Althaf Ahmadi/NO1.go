package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fx float32

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&fx)

	x := 2/(fx+5) - 5

	strX := fmt.Sprintf("%.0f", x) 
	xInt, err := strconv.Atoi(strX)
	if err != nil {
		fmt.Println("Error konversi:", err)
		return
	}

	fmt.Println("Nilai x (integer):", xInt)
}
