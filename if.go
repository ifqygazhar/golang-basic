package main

import "fmt"

func main() {
	name := "ifqy"

	if name == "ifqy" {
		fmt.Println("halo ifqy")
	} else if name == "mariam" {
		fmt.Println("halo mariam")
	} else {
		fmt.Println("boleh kenalan??")
	}

	if panjang := len(name); panjang > 5 {
		fmt.Println("sudah benar")
	} else {
		fmt.Println("terlalu panjang")
	}
}
