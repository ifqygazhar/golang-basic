package main

import "fmt"

type Address struct {
	daerah, provinsi, negara string
}

func main() {
	adress1 := Address{
		daerah:   "pangandaran",
		provinsi: "jabar",
		negara:   "indonesia",
	}

	var adress3 *Address = &adress1

	var adress5 *Address = &Address{"cimahi", "bandung", "indonesia"}

	adress2 := &adress1

	adress2.daerah = "bandung"

	*adress2 = Address{"jakarta", "jabar", "indonesia"}

	fmt.Println("ke satu :", adress1)
	fmt.Println("ke dua :", adress2)
	fmt.Println("ke tiga :", adress3)

	var adress4 *Address = new(Address)
	adress4.daerah = "Cimerak"
	adress4.provinsi = "jabar"
	adress4.negara = "indonesia"
	fmt.Println(adress4)
	fmt.Println(adress5)
}
