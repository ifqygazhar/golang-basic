package main

import "fmt"

type hashName interface {
	Getname() string
}

func sayHello(hashname hashName) {
	fmt.Println("halo", hashname.Getname())
}

type Person struct {
	Name string
}

func (person Person) Getname() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) Getname() string {
	return animal.Name
}

func main() {
	var az Person
	az.Name = "ifqy"
	sayHello(az)

	cat := Animal{Name: "kucing"}
	sayHello(cat)
}
