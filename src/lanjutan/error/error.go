/*
Tipe error di Go adalah sebuah interface bawaan yang memiliki satu metode.
Metode Error() mengembalikan representasi string dari error. Anda dapat
membuat tipe error Anda sendiri dengan mengimplementasikan metode Error().
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

// built-n error interface
type error interface {
	Error() string
}

type ErrorGw struct {
	Kapan time.Time
	Apaan string
}

func (e ErrorGw) Error() string {
	return fmt.Sprintf("pas %v, %s", e.Kapan, e.Apaan)
}

func jalankan() error {
	return &ErrorGw{
		time.Now(),
		"gagal :v",
	}
}

func main() {
	// Fungsi sering mengembalikan nilai error, menangani error dilakukan dengan
	// memeriksa apakah error tersebut == nil.
	angka, err := strconv.Atoi("430")

	// error == nil -> sukses
	// error != nil -> gagal
	if err != nil {
		fmt.Println("gagal konversi string ke angka")
		return
	}

	fmt.Println(angka)

	if err := jalankan(); err != nil {
		fmt.Println(err)
	}
}
