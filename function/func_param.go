package main

import "fmt"

type Filter func(string) string

func sayHellloFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("helo", nameFilter)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHellloFilter("azhar", spamFilter)
	sayHellloFilter("anjing", spamFilter)
}
