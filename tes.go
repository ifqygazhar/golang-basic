package main

import (
	"fmt"
)

func main() {

	var i int
	var b int
	fmt.Println("Enter an integer value : ")

	_, err := fmt.Scanf("%d", &i)
	fmt.Println("Enter an integer value : ")
	_, be := fmt.Scanf("%d", &b)

	if err != nil {
		fmt.Println(err)
	}
	if be != nil {
		fmt.Println(be)
	}

	fmt.Println("You have entered : ", i)
	fmt.Println("You have entered : ", b)
	fmt.Println("result i * b = ", i*b)

}
