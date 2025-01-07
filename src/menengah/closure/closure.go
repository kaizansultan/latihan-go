/*
Closure dalam Go adalah fungsi yang "mengingat" variabel di luar tubuhnya yang
dirujuk selama pembuatannya. Closure memungkinkan fungsi memiliki konteks
uniknya sendiri.
*/
package main

import "fmt"

func hitung() func() int {
	hitungan := 0

	return func() int {
		hitungan++
		return hitungan
	}
}

func main() {
	c1 := hitung() // membuat closure pertama
	c2 := hitung() // membuat closure kedua

	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c1())
	fmt.Println(c2())

	// Masing-masing closure membuat instance baru dari hitungan dengan nilai 0
	// sehingga nilainya tidak tercampur antar closure
}
