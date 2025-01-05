/*
Struct adalah sekumpulan dari field.
*/
package main

import "fmt"

type Orang struct {
	nama    string
	panjang uint8
}

func main() {
	fmt.Println(Orang{"Izka", 15})

	orangA := Orang{"Azen", 20}

	// diakses menggunakan titik
	fmt.Println(orangA.nama)
	fmt.Println(orangA.panjang)

	/*
		Field bisa diakses dengan pointer. Misal untuk mengakes field panjang
		-> (*orangA).panjang

		Notasi di atas merepotkan, jadi Anda hanya perlu menulisnya dengan
		-> orangA.panjang
		tanpa dereferencing eksplisit
	*/
	orangA.panjang = 21
	fmt.Println(orangA)
}
