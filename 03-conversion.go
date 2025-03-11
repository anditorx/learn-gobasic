package main

import "fmt"

/*
Untuk melakukan print / console.log ,
kita perlu melakukan import module fmt terlebih dahulu
*/

func main() {
	var name = "Andito Nasir"
	// var e = name[0]
	var eString = string(name[0])

	// fmt.Println(e)
	fmt.Println(eString)
}
