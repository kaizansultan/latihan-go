package main

import "fmt"

const Pi = 3.14

/*
Aturan Konstanta
- `const` tidak bisa dideklarasi menggunakan `:=`.
- Nilainya harus ditentukan saat kompilasi.
- Anotasi tipe tidak diwajibkanm, namun harus diketahui pada kompilasi.
- Nilai tidak dapat diubah.

Konvensi Penamaan Identifier
- gunakan PascalCase untuk konstanta yang diekspor,
- gunakan camelCase untuk konstanta lokal,
- gunakan uppercase SNAKE_CASE untuk konstanta global
*/

func main() {
	fmt.Println(100 * Pi)
}
