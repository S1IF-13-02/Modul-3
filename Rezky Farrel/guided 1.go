package main

import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan nilai sisi: ")
	fmt.Scan(&sisi)

	volumeKubus := sisi * sisi * sisi
	fmt.Printf("Volume kubus dengan sisi %d adalah %d\n", sisi, volumeKubus)
}