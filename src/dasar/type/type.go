package main

import "fmt"

// Deklarasi variable bisa difaktor ke dalam blok, seperti import.
var (
	nama      string  = "Mauza"
	umur      int8    = 20
	bb        float32 = 55.6
	berbatang bool    = true

	panjang byte = 16
)

/*
Tipe-tipe Dasar dalam Go:

bool string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

float32 float64
complex 64 complex 128

byte => uint8
rune => int32


int, uint, uintptr adalah 32 bits pada 32-bit sys dan 64 bits pada 64-bit sys.

Selalu gunakan `int` untuk angka, kecuali Anda memiliki alasan tertentu.
*/

func main() {
	fmt.Println(nama, umur, bb, berbatang)
}
