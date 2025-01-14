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
	Anu()
}

/*
Implementasi interface dalam Go bersifat implisit. Tidak ada kata kunci seperti
implements atau extends di Go. Sebuah tipe dikatakan mengimplementasikan
interface jika tipe tersebut memiliki metode dengan tanda tangan (signature)
yang sama dengan metode yang dideklarasikan pada interface.
*/

// Tipe Hewan mengimplementasikan interface Makhluk karena memiliki metode
// Makan(string) string.
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

func (o *Orang) Anu() {
	fmt.Println(o.Anu)
}

func (h Hewan) Anu() {
	fmt.Println(h.Anu)
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

	// Dalam Go, implementasi tidak dideklarasikan secara eksplisit.
	// Nilai dari interface adalah pasangan tuple (value, type).
	anu.Anu()

	var goib Makhluk // nil (tidak memiliki value dan type)
	fmt.Printf("Interface goib: (%v, %T)\n", goib, goib)
	// goib.Anu()
	// Runtime error: invalid memory address or nil pointer dereference

	// Empty interface  adalah interface yang tidak memiliki metode,
	// sehingga dapat menampung nilai dari tipe apa saja.
	var any interface{}
	describe(any)

	any = 30
	describe(any)
	// tipe := any.(string)
	// ^ panic
	tipe := any.(int)
	fmt.Println(tipe)

	any = "Neptunus"
	describe(any)
	nilai, ok := any.(int)
	// ^ tidak panik jika diekstrak begini
	fmt.Println(nilai, ok)
	switchType(any)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
