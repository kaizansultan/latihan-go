# Konstanta di Go

Dalam Go, **konstanta** digunakan untuk mendefinisikan nilai yang tidak berubah selama waktu eksekusi program. Konstanta memungkinkan Anda memberi nama pada nilai tetap, yang meningkatkan keterbacaan dan pemeliharaan kode.

## Sintaksis Dasar

Konstanta didefinisikan menggunakan kata kunci `const`. Berikut adalah sintaksisnya:

```go
const identifier = value
```

-   **identifier**: Nama konstanta.
-   **value**: Nilai tetap yang diberikan.

## Aturan Konstanta

1. Nilainya **harus ditentukan pada saat kompilasi**.
    - Contoh yang valid: `const x = 42`
    - Contoh yang tidak valid: `const x = runtimeValue` (karena `runtimeValue` dihitung saat runtime).
2. **Tipe opsional**:

    - Jika tidak diberi tipe, tipe akan ditentukan berdasarkan konteks penggunaannya.
    - Contoh:
        ```go
        const x = 10        // Tidak ditentukan tipe (default int)
        const y float64 = 3.14 // Ditentukan tipe float64
        ```

3. Konstanta **tidak dapat diubah** setelah didefinisikan.

4. Nilai konstanta harus berupa tipe:
    - Bilangan (integer, floating-point).
    - Boolean.
    - String.
    - Karakteristik lainnya seperti kompleks atau runtime value tidak didukung.

## Contoh

### Konstanta Sederhana

```go
package main

import "fmt"

func main() {
    const Pi = 3.14
    const Name = "GoLang"
    fmt.Println("Pi:", Pi)
    fmt.Println("Name:", Name)
}
```

### Konstanta dengan Tipe

```go
const x int = 42
const y float64 = 3.14
```

## Konstanta Iota

`iota` adalah mekanisme bawaan Go untuk menghasilkan nilai berturut-turut secara otomatis dalam deklarasi konstanta.

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)

const (
    _ = iota     // 0 (diabaikan dengan `_`)
    KB = 1 << (10 * iota) // 1024
    MB                  // 1048576
    GB                  // 1073741824
)
```

-   **Penjelasan**:
    -   `iota` dimulai dari `0` dan bertambah `1` untuk setiap baris dalam satu blok konstanta.
    -   Sangat berguna untuk membuat konstanta yang berurutan atau bit-shifting.

## Keuntungan Menggunakan Konstanta

1. **Keamanan**: Mencegah perubahan nilai yang tidak disengaja.
2. **Keterbacaan**: Nama konstanta memberikan makna pada nilai.
3. **Optimisasi**: Compiler dapat mengoptimalkan konstanta lebih baik daripada variabel.

## Perbedaan Konstanta dan Variabel

| **Aspek**             | **Konstanta**                  | **Variabel**                |
| --------------------- | ------------------------------ | --------------------------- |
| Dideklarasikan dengan | `const`                        | `var`                       |
| Nilai dapat berubah?  | Tidak                          | Ya                          |
| Dihitung saat?        | Kompilasi                      | Runtime                     |
| Tipe?                 | Harus diketahui pada kompilasi | Bisa berubah selama runtime |
