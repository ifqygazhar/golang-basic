package main

import "fmt"

func getHello(name string, name2 string) string {
	if name == name2 {
		return "namanya " + name + name2
	} else {
		return "namanya tidak sama"
	}
}

func main() {
	result := getHello("ifqy", "azhar")
	fmt.Println(result)

}
