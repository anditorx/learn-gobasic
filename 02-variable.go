package main

import "fmt"

/*
Untuk melakukan print / console.log ,
kita perlu melakukan import module fmt terlebih dahulu
*/

func main() {
	var name string
	name = "Andi"
	fmt.Println(name)

	name = "Budi"
	fmt.Println(name)

	studentName := "Dante"
	fmt.Println(studentName)

	studentName = "Saras"
	fmt.Println(studentName)

	// NOTE: Multiple variable
	var (
		firstName = "Putra"
		midleName = "Dirga"
		lastName  = "Djayadi"
	)
	fmt.Println(firstName, midleName, lastName)
}
