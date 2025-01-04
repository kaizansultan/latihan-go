package main

import "fmt"

func sapa(jam uint8) string {
	if jam > 0 && jam < 11 {
		return "Selamat pagi"
	} else if jam >= 11 && jam < 17 {
		return "Selamat siang"
	} else if jam >= 17 && jam < 20 {
		return "Selamat sore"
	} else {
		return "Selamat malam"
	}
}

func main() {
	fmt.Println(sapa(13))

	// Untuk efesiensi dan menghindari pencemaran variabel global, kita bisa
	// membuat statement secara langsung setelah `if`, seperti pada `for`.
	if panjang := len(sapa(22)); panjang > 12 {
		fmt.Println("Sapaan lebih dari 12 huruf 😙.", panjang)
	} else {
		fmt.Println("Sapaan kurang dari 12 huruf 😡.", panjang)
	}
}
