package main

import (
	"fmt"
)

func main() {
	//farenheit ke celcius & celcius ke k
	print("masukan angka konversi farenheit ke celcius  :")
	var a int
	fmt.Scan(&a)
	var b int = 32
	var c int = 5 / 9

	hasil := a - b
	hasil2 := hasil * c

	fmt.Println("hasil f ke c:", hasil2, "c")

}
