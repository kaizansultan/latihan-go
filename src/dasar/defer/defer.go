/*
Defer dieksekusi terbalik dari urutan deklarasi, tidak peduli apakah fungsi
utama berhasil atau error. Deferred function calls disimpan dalam stack. Defer
berguna misalnya untuk menutup file, melepas resource, operasi cleanup ketika
menangani panic.

Perhatian: defer dapat memiliki overhead kecil dalam hal performa, terutama
jika digunakan berulang kali dalam loop. Oleh karena itu, lebih bijaksana
untuk tidak menggunakannya dalam situasi yang sangat membutuhkan kecepatan
tinggi.
*/
package main

import "fmt"

func main() {
	// Menunda pemanggilan fungsi
	defer fmt.Println("fungsi 1")

	fmt.Println("fungsi 2")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	// Menunda pemanggilan fungsi lainnya
	defer fmt.Println("fungsi 3")

	fmt.Println("fungsi 4")
}
