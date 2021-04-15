package main

import (
	"fmt"
)

func main() {
	type anjay string
	type nikah bool

	var aku anjay = "ifqy"
	var dia anjay = "Mariam"

	var kawin nikah = false

	fmt.Println("Saya", aku, "dan", dia, "nikah =", kawin)

}
