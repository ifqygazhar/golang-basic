package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "ifqy",
		"address": "pangandaran",
	}
	person["gf"] = "Mariam"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["gf"])

	book := make(map[string]string)
	book["title"] = "buku go-lang"
	book["author"] = "ifqy gifha azhar"
	book["wrong"] = "salah"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))

	nomor := make(map[string]int)
	nomor["satu"] = 1
	nomor["dua"] = 2

	fmt.Println(len(nomor))

}
