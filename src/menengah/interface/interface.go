/*
Interface dalam Go adalah tipe data yang didefinisikan sebagai sekumpulan method
signatures. Interface tidak memiliki implementasi dari metode-metode tersebut,
melainkan hanya mendeklarasikan kontrak bahwa tipe tertentu dapat dikatakan
"mengimplementasikan" interface jika tipe tersebut memiliki semua metode yang
sesuai dengan deklarasi pada interface tersebut.
*/
package main

import "fmt"

type Makhluk interface {
	Makan(string) string
}

/*
Implementasi interface dalam Go bersifat implisit. Tidak ada kata kunci seperti
implements atau extends di Go. Sebuah tipe dikatakan mengimplementasikan
interface jika tipe tersebut memiliki metode dengan tanda tangan (signature)
yang sama dengan metode yang dideklarasikan pada interface.
*/

// Tipe Hewan mengimplementasikan interface Makhluk karena memiliki metode Makan(string) string.
type Hewan string

type Orang struct {
	nama string
	bb   uint32
}

func (h Hewan) Makan(m string) string {
	return string(h) + " makan " + m
}

func (o *Orang) Makan(m string) string {
	return o.nama + " makan " + m
}

func main() {
	var anu Makhluk
	kucing := Hewan("kucing")
	dandy := Orang{"Dandy Pangalima", 60}

	// anu Hewan mengimplementasikan Makhluk
	anu = kucing
	fmt.Println(anu.Makan("ikan"))

	// anu *Orang mengimplementasikan Makhluk
	anu = &dandy
	fmt.Println(anu.Makan("ikan"))

	// Kode di bawah error karena bukan *Orang
	// dan tidak mengimplementasikan Makhluk
	// anu = dandy
}
