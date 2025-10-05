package main
import "fmt"

func main() {
var r int
const PI float64 = 3.1415926535

fmt.Print("Masukkan jari-jari bola: ")
fmt.Scan(&r)
volume := (4.0 / 3.0) * PI * float64(r) * float64(r) *float64(r)
luas := 4 * PI * float64(r) * float64(r)
fmt.Printf ("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}