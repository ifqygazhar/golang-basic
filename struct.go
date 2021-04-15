package main

import "fmt"

type Customer struct {
	name, address string
	age           int
	married       bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("helo", name, "nama saya adalah", customer.name)
}

func (a Customer) sayHu() {
	fmt.Println("huu from", a.name)
}

func main() {
	var ifqy Customer
	ifqy.name = "ifqy gifha azhar"
	ifqy.address = "cimerak"
	ifqy.age = 16
	ifqy.married = false

	ifqy.sayHi("mariam")
	ifqy.sayHu()
	// mariam := Customer{
	// 	name:    "mariam saputri",
	// 	address: "cimahi",
	// 	age:     16,
	// 	married: false,
	// }
	// fmt.Println(mariam)

	// fmt.Println(ifqy)
	// fmt.Println(ifqy.name)
	// fmt.Println(ifqy.address)
	// fmt.Println(ifqy.age)
}
