package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "halo bre"
	} else {
		return "halo " + name
	}
}

func main() {
	result := sayHello("eko")
	fmt.Println(result)

	fmt.Println(sayHello(""))
}
