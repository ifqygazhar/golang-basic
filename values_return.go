package main

import "fmt"

func getFullname() (firstname, midlename, lastname string) {
	firstname = "ifqy"
	midlename = "gifha"
	lastname = "azhar"

	return
}

func main() {
	a, b, c := getFullname()
	fmt.Println(a, b, c)
}
