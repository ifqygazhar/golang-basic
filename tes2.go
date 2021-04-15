package main

import "fmt"

type Alamat struct {
	Daerah, Kota, Negara string
}

func tesAlamat(alamat *Alamat) {
	alamat.Negara = "indonesia"
}

func main() {
	address := Alamat{"bandung", "jabar", "belanda"}
	var baru *Alamat = &address
	tesAlamat(baru)
	fmt.Println(baru)
}
