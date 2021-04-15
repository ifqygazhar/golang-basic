package main

import "fmt"

func main() {

	for counter := 1; counter < 10; counter++ {
		fmt.Println("perulangan ke = ", counter)

	}
	slice := []string{"ifqy", "gifha", "azhar"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, name := range slice {
		fmt.Println("name =", name)
	}

	person := make(map[string]string)
	person["nama"] = "ifqy"
	person["title"] = "programer"

	for _, value := range person {
		fmt.Println(value)
	}
}
