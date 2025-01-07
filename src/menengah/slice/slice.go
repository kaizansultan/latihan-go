/*
Jika array ukurannya tetap, slice tidak. Dalam praktiknya, slice lebih umum
dijumpai daripada array.
*/
package main

import "fmt"

func main() {
	urutan := [7]int{1, 2, 4, 5, 8, 10, 13}

	// slice[low:high] -> memilih elemen low hingga sebelum high
	var terpilih []int = urutan[1:4]
	fmt.Println(terpilih)

	/*
		Slice tidak menyimpan data, ia hanya mendeskripsikan bagian mana dari
		array yang ingin diindex.

		Karena slice adalah referensi ke array, jika kita mengubah elemen dalam
		slice, maka elemen tersebut akan mengubah nilai yang ada dalam array
		yang mendasarinya.

		Jika ada slice lain yang juga merujuk ke array yang sama, perubahan
		yang dilakukan pada salah satu slice akan terlihat pada slice lainnya.
	*/

	awal := urutan[0:4]
	tengah := urutan[2:6]

	fmt.Println(awal)
	fmt.Println(tengah)

	awal[3] = 88

	fmt.Println(urutan)
	fmt.Println(awal)
	fmt.Println(tengah)

	// slice literals:
	jawaban := []bool{true, true, false}
	fmt.Println(jawaban)

	sus := []struct {
		nama string
		sus  bool
	}{
		{"Amar", false},
		{"Tole", true},
		{"Sponge", true},
		{"Roro", false},
		{"Edna", true},
	}
	fmt.Println(sus)

	// Anda bisa membuang low atau high.
	fmt.Println(urutan[0:5])
	fmt.Println(urutan[:5])

	fmt.Println(len(awal))
	fmt.Println(cap(awal))

	// Memperluas slice hingga kapasitas maksimumnya. Jika Anda melampauinya,
	// akan terjadi runtime error.
	awal = awal[:7]
	fmt.Println(len(awal))

	// Jika slice kosong, maka nilainya adalah nil.
	var kosong []int
	fmt.Println(kosong, len(kosong), cap(kosong), kosong == nil)

	// Slice bisa mengandung tipe apapun, termasuk slice lainnya.
	nama := [][]string{
		{"Kaizan", "Sultan"},
		{"Zarul", "Sultan"},
		{"Zay", "Majid"},
	}
	fmt.Println(nama)

	// append
	printSlice(kosong)
	kosong = append(kosong, 0)
	printSlice(kosong)
	kosong = append(kosong, 1)
	printSlice(kosong)
	kosong = append(kosong, 2, 3, 4)
	printSlice(kosong)

	/*
		Slice dapat dibuat secara dinamis menggunakan fungsi bawaan make.
		Fungsi ini akan mengalokasikan array kosong dan mengembalikan slice
		yang merujuk ke array tersebut. Anda dapat menentukan panjang slice
		dengan dua argumen (make([]T, length)) atau menambahkan kapasitas
		dengan argumen ketiga (make([]T, length, capacity)). Panjang slice
		dapat diubah dengan re-slicing selama tidak melebihi kapasitasnya.
		Contoh: make([]int, 0, 5) menghasilkan slice dengan panjang 0 dan
		kapasitas 5, yang dapat diperpanjang hingga kapasitas maksimum
		dengan re-slicing.
	*/
	kursi := make([]int, 2, 5)
	printSlice(kursi)
	kursiDiambil := kursi[:2]
	printSlice(kursiDiambil)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
