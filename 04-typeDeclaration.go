package main

import "fmt"

/* NOTE: Type Declaration
Type Declaration adalah kemampuan untuk membuat ulang sebuah type yang sudah ada.
Biasanya digunakan untuk membuat type yang lebih spesifik dari type yang sudah ada.
*/

func main() {
	type NoKTP string
	var ktpEko NoKTP = "1111111111"

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("2222222222"))
}
