package main

import "fmt"

func main() {
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// month[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei baru"
	// fmt.Println(slice1)

	var slice2 = month[11:]
	fmt.Println("slice2",slice2)

	var slice3 = append(slice2, "Azhar")
	fmt.Println("slice3",slice3)

	slice3[1] = "Desember baru"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println("month",month)
	fmt.Println("cap month = ",cap(month))
	fmt.Println("len month = ",len(month))

	newSlice := make([]string,2,5)
	newSlice[0] = "senin"
	newSlice[1] = "selasa"
	

	fmt.Println(newSlice)

	fmt.Println("len = ",len(newSlice))
	fmt.Println("cap = ",cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

}
