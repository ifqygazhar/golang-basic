package main

import "fmt"

func main() {
	// belanja di toko
	var harga int
	fmt.Print("masukan nama barang = ")
	var barang string
	fmt.Scan(&barang)
	fmt.Print("bayaran uang anda = ")
	var uang int
	fmt.Scan(&uang)

	switch barang {
	case "Laptop":
		harga = 10000
	case "Handphone":
		harga = 5000
	default:
		fmt.Println("tidak ada barang", barang)

	}
	if harga == 10000 {
		fmt.Println("terimakasih sudah berbelanja, sisa uang anda = ", uang-harga)
	} else if harga == 5000 {
		fmt.Println("terimakasih sudah berbelanja, sisa uang anda = ", uang-harga)
	} else {
		fmt.Println("uang tidak cukup / barang", barang, "tidak ada")
	}
}
