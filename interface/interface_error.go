package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("dibagi NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("hasilnya :", hasil)
	} else {
		fmt.Println("error nya :", err.Error())
	}
}
