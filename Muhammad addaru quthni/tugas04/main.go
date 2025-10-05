package main

import (
    "fmt"
)

func celsiusToFahrenheit(c float64) float64 {
    return (c * 9.0 / 5.0) + 32
}

func celsiusToReaumur(c float64) float64 {
    return c * 4/5
}

func celsiusToKelvin(c float64) float64 {
    return c + 273.15
}

func main() {
    var f float64
    fmt.Print("Masukkan suhu Temperatur Celcius: ")
    fmt.Scanln(&f)

    c := celsiusToFahrenheit(f)
    r := celsiusToReaumur(f)
    k := celsiusToKelvin(f)

    fmt.Printf("Temperature di Reaumur: %.f\n", r)
    fmt.Printf("Temperature di Fahrenheit: %.f\n", c)
    fmt.Printf("Temperature di Kelvin: %.f\n", k)
}