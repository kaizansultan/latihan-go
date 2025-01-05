package main

import "fmt"

func main() {
	// Panjang array adalah bagian dari tipenya,
	// jadi tidak bisa diubah ukurannya.
	var daftarPanjang = [4]int{12, 20, 18, 15}
	fmt.Println(daftarPanjang)

	var sapa [3]string
	sapa[0] = "Halo"
	sapa[1] = "Selamat datang"
	fmt.Println(sapa)

	sapa[2] = "Siapa lu?"
	fmt.Println(sapa)
	fmt.Println(sapa[1])

	daftarNilai := [8]int8{}
	fmt.Println(daftarNilai)
}
