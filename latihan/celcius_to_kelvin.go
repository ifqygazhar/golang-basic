package main

import (
	"fmt"
	"math"
)

func main() {
	//celcius ke kelvin
	print("masukan angka konversi celcius ke kelvin  :")
	var a int
	fmt.Scan(&a)

	var d float64 = 273.15
	e := int(math.Round(d)) // konversi float ke integer

	hasil := a + e

	fmt.Println("hasil c ke k :", hasil, "k")
}
