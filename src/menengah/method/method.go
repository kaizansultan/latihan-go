/*
Go tidak punya class, tapi Anda bisa membuat method untuk type. Method adalah
sebuah fungsi dengan argumen receiver spesial. Method hanyalah  fungsi biasa
dengan tambahan receiver argument.
*/
package main

import "fmt"

type Orang struct {
	nama string
	bb   uint32
}

func (o Orang) Makan() {
	fmt.Println(o.nama + " lagi makan")
}

func Makan(o Orang) {
	fmt.Println(o.nama + " lagi makan")
}

func (o Orang) AmbilBB() uint32 {
	return o.bb
}

// value receiver
func (o Orang) NaikBB(n uint32) {
	o.bb = o.bb + n
}

// pointer receiver
func (o *Orang) NaikanBB(n uint32) {
	o.bb = o.bb + n
}

// pointer receiver fungsi biasa
func KurangiBB(o *Orang, n uint32) {
	o.bb = o.bb - n
}

// Tidak harus ke type struct. Bisa ke semua type apa aja.
type AngkaGuwe int32

func (x AngkaGuwe) Tambahkan(y AngkaGuwe) AngkaGuwe {
	return x + y
}

func main() {
	agus := Orang{"Agus", 80}
	agus.Makan()
	fmt.Println(agus.AmbilBB())

	// Makan bisa ditulis sebagai fungsi biasa tanpa perubahan pada fungsinya.
	Makan(Orang{"Deni", 97})

	angka := AngkaGuwe(4)
	fmt.Println(angka.Tambahkan(5))

	agus.NaikBB(70)
	fmt.Println(agus.AmbilBB())

	// Gunakan pointer receiver untuk mengubah nilai asli. Sangat efesien,
	// karena tidak perlu membuat salinan struct, terutama jika ukurannya besar
	agus.NaikanBB(70)
	// diinterpretasikan -> (&agus).NaikanBB(70)
	fmt.Println(agus.AmbilBB())

	KurangiBB(&agus, 50)
	// KurangiBB(agus, 50) -> tidak otomatis diinterpretasikan
	fmt.Println(agus.AmbilBB())

}
