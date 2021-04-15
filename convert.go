package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 10000
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	name := "ifqy"
	e := name[0]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
