package main

import "fmt"

type Address struct {
	Daerah, Kota, Negara string
}

func ChangeCountryAddress(address *Address) { //pointer
	address.Negara = "Indonesia"
}

func main() {
	address := Address{"Jakarta", "DKI jkt", "Belanda"}
	var alamat *Address = &address // & = ngambil pointer
	ChangeCountryAddress(alamat)

	fmt.Println(alamat) //print pointer (data keubah)
}
