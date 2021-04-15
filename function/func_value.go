package main

import "fmt"

func sayGoodbye(name string) string {
	return "goodbye " + name
}

func main() {
	goodbye := sayGoodbye
	result := goodbye("azhar")
	fmt.Println(result)
}
