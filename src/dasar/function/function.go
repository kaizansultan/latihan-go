package main

import "fmt"

func jumlahkan(x int, y int) int {
	return x + y
}

// Jika dua atau lebih parameter bernama yang memiliki tipe sama, Anda bisa
// membuang anotasi tipenya kecuali yang paling akhir
func kurangkan(x, y, z int) int {
	return x - y - z
}

// Fungsi bisa mengembalikan lebih dari dua nilai
func tukar(a, b string) (string, string) {
	return b, a
}

// Nilai return bisa diberi nama

func main() {
	fmt.Println(jumlahkan(2, 6))
	fmt.Println(kurangkan(10, 2, 4))

	a, b := tukar("Halo", "Guys")
	fmt.Println(a, b)
}
