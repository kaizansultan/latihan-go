/*
Dalam Go, range adalah bentuk khusus dari for loop yang digunakan untuk
mengiterasi elemen dalam slice atau map. Ketika digunakan pada slice, range
menghasilkan dua nilai pada setiap iterasi: indeks elemen dan salinan dari
elemen pada indeks tersebut.
*/
package main

import "fmt"

func main() {
	nomorBalap := []int{46, 28, 19, 54, 44, 90}
	for index, nomor := range nomorBalap {
		fmt.Println(index, nomor)
	}

	// Jika hanya ingin mengambil nilai:
	for _, nomor := range nomorBalap {
		fmt.Println(nomor)
	}

	// Jika hanya ingin mengambil index:
	for index, _ := range nomorBalap {
		fmt.Println(index)
	}

	// Atau lebih singkatnya:
	for index := range nomorBalap {
		fmt.Println(index)
	}
}
