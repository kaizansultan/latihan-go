package main

import "fmt"

func main() {
	harga, diskon := 10000, 2000

	pointerA := &harga
	// membaca harga lewat pointerA
	fmt.Println(*pointerA)
	fmt.Println(pointerA)

	// mengubah harga lewat pointerA
	// menggunakan "dereferencing"
	*pointerA += 3000
	fmt.Println(harga)

	// bedakan jika tidak pake pointer
	pointerB := diskon
	pointerB /= 2
	fmt.Println(diskon)

	// menggunakan var
	var pointerC *int = &diskon
	*pointerC = *pointerC * 3
	fmt.Println(diskon)
}
