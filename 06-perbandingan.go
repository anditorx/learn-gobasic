package main

import "fmt"

/* NOTE: Operasi Perbandingan
Operasi perbandingan adalah operasi untuk membandingkan dua buah data
Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
Jika hasil operasinya adalah benar, maka nilainya adalah true
Jika hasil operasinya adalah salah, maka nilainya adalah false

*/

func main() {
	var name1 = "Budi"
	var name2 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)
}
