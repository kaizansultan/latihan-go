package main

import "fmt"

var mobil, buah string
var berbatang bool
var umur int

// Deklarasi var bias memiliki initializer, satu variabel satu initializer.
var namaDepan, namaBelakang string = "Kaizan", "Sultan"

// Jika initializer ada, anotasi tipe bisa dibuang.
var merk, usia, hidup = "Ahamay", 8, true

func main() {
	var i string
	fmt.Println(i, mobil, buah, berbatang, umur)
	/*
		string kosong -> ""
		bool kosong -> false
		int kosong -> 0
	*/

	fmt.Println(namaDepan, namaBelakang, merk, usia, hidup)

	// Deklarasi variabel bisa dipersingkat menggunakan operator `:=`. Namun
	// hanya bisa dilakukan di dalam fungsi. Karena di luar fungsi, setiap
	// statement harus diawali dengan kata kunci (misalnya `var`, `func`, dll).
	planet := "Jupiter"
	fmt.Println("Aku tinggal di planet", planet)
}
