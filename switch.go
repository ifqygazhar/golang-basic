package main

import "fmt"

func main() {
	var name string
	fmt.Print("masukan nama = ")
	fmt.Scan(&name)

	switch name {
	case "ifqy":
		fmt.Println("halo ifqy")
	case "joko":
		fmt.Println("halo joko")
	default:
		fmt.Println("tidak ada dalam daftar")
	}

	// switch panjang := len(name); panjang > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")

	// }
	// fmt.Println(len(name))

	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("panjang")
	case panjang > 5:
		fmt.Println("lumayan")
	default:
		fmt.Println("terlalu panjang")
	}
}
