# Pointer ke Struct

Teknik menggunakan pointer ke `struct` seperti `p = &Vertex{1, 2}` berguna ketika kita ingin **mengakses dan memodifikasi data di memori secara langsung**, tanpa membuat salinan data. Pointer juga membantu dalam **efisiensi memori** saat bekerja dengan data besar, karena kita hanya menyimpan referensi ke data, bukan datanya sendiri.

Berikut adalah contoh nyata yang lebih mudah dipahami:

## Studi Kasus: Sistem Pemesanan Tiket Pesawat

Kita memiliki sistem untuk memesan tiket pesawat, di mana setiap pengguna memiliki detail pemesanan. Pointer digunakan agar kita bisa **mengupdate data pengguna secara langsung** tanpa menduplikasi data.

```go
package main

import "fmt"

// Struct untuk menyimpan data penumpang
type Passenger struct {
	Name  string
	Email string
	Seat  string
}

// Fungsi untuk mengupdate kursi penumpang
func UpdateSeat(p *Passenger, newSeat string) {
	p.Seat = newSeat // Modifikasi data langsung melalui pointer
}

func main() {
	// Membuat instance Passenger
	passenger := &Passenger{
		Name:  "Kaizan Sultan",
		Email: "kaizan@example.com",
		Seat:  "12A",
	}

	// Menampilkan data awal
	fmt.Println("Before Update:", *passenger)

	// Memperbarui kursi menggunakan fungsi dengan pointer
	UpdateSeat(passenger, "14C")

	// Menampilkan data setelah update
	fmt.Println("After Update:", *passenger)
}
```

## Penjelasan

1. **Pointer `passenger`:**

    ```go
    passenger := &Passenger{
        Name:  "Kaizan Sultan",
        Email: "kaizan@example.com",
        Seat:  "12A",
    }
    ```

    - Di sini, kita membuat pointer ke `Passenger`. Dengan menggunakan `&Passenger{}`, kita menyimpan alamat memori dari struct `Passenger` di variabel `passenger`.

2. **Mengupdate Data Melalui Pointer:**

    ```go
    func UpdateSeat(p *Passenger, newSeat string) {
        p.Seat = newSeat // Modifikasi langsung pada data asli
    }
    ```

    - Fungsi ini menerima pointer ke `Passenger` (`*Passenger`) sehingga perubahan yang dilakukan pada `p.Seat` akan langsung mengupdate data asli tanpa membuat salinan.

3. **Efisiensi dan Kepraktisan:**
    - Jika tidak menggunakan pointer, kita harus mengembalikan struct yang telah diubah, seperti:
        ```go
        func UpdateSeat(p Passenger, newSeat string) Passenger {
            p.Seat = newSeat
            return p
        }
        ```
    - Pendekatan ini kurang efisien karena melibatkan salinan struct setiap kali ada perubahan.

## Output Program

```console
Before Update: {Kaizan Sultan kaizan@example.com 12A}
After Update: {Kaizan Sultan kaizan@example.com 14C}
```

## Kapan Harus Menggunakan Pointer?

1. **Data Besar:** Struct dengan banyak field atau data yang membutuhkan efisiensi memori.
2. **Perubahan Data:** Ketika fungsi perlu **mengubah** data asli dan perubahan itu harus terlihat di luar fungsi.
3. **Referensi Tunggal:** Ketika banyak bagian program membutuhkan referensi ke data yang sama.

Teknik ini sangat sering digunakan dalam sistem seperti **CRUD database**, **API handler**, dan **algoritma berbasis objek**.

> Nah, kalo begitu kenapa ga semua pembuatan instanse dari struct menggunakan pointer?

## Kenapa Tidak Semuanya Menggunakan Pointer?

Tidak semua instansi struct menggunakan pointer karena **pointer memiliki trade-off** dan penggunaannya bergantung pada kebutuhan spesifik program. Berikut adalah alasan mengapa tidak semua instansi struct menggunakan pointer:

### 1. Kebutuhan Mutasi vs. Non-Mutasi

-   **Pointer digunakan** saat Anda ingin **mengubah data asli** (mutasi) atau membutuhkan efisiensi memori.
-   **Struct langsung digunakan** saat Anda tidak perlu mengubah data atau jika hanya bekerja dengan salinan (immutable). Hal ini lebih aman dari bug karena data asli tidak bisa diubah tanpa sengaja.

**Contoh:**
Jika Anda hanya membaca data tanpa mengubahnya:

```go
user := User{Name: "Kaizan", Age: 22} // Tidak perlu pointer
fmt.Println(user.Name) // Operasi baca saja
```

### 2. Sederhana dan Mudah Dimengerti

-   Struct langsung lebih mudah dipahami karena tidak melibatkan dereferensi (`*`) atau operator pointer (`&`).
-   Cocok untuk kode sederhana atau situasi di mana modifikasi data tidak diperlukan.

**Contoh:**

```go
type Point struct {
    X, Y int
}

p := Point{3, 4} // Struct langsung
fmt.Println(p.X, p.Y) // Mudah diakses
```

### 3. Efisiensi pada Data Kecil

-   Struct langsung lebih efisien jika ukuran struct kecil (misalnya beberapa field integer). Biaya salinan kecil dibandingkan dengan biaya tambahan dereferensi pointer.
-   Pointer menambahkan overhead seperti:
    -   Penyimpanan alamat di memori.
    -   Dereferensi saat membaca atau menulis data.

**Contoh:**

```go
type SmallStruct struct {
    A, B int
}

s1 := SmallStruct{1, 2} // Salinan lebih murah
s2 := s1                // Salin langsung
fmt.Println(s2)
```

### 4. Garbage Collection (GC)

-   Pointer membuat data dialokasikan di heap (memori dinamis). Data ini harus dilacak oleh garbage collector.
-   Struct langsung disimpan di stack (memori lokal), yang lebih cepat dan tidak membebani GC.

**Perbedaan:**

-   Dengan pointer:
    ```go
    p := &SmallStruct{1, 2} // Dialokasikan di heap
    ```
-   Tanpa pointer:
    ```go
    s := SmallStruct{1, 2} // Dialokasikan di stack
    ```

### 5. Konvensi Pemrograman

-   Pointer biasanya digunakan untuk **data besar atau kompleks**, sedangkan struct langsung digunakan untuk **data kecil atau sederhana**.
-   Ini adalah praktik umum untuk menjaga **keseimbangan antara performa dan keterbacaan kode**.

## Kapan Menggunakan Masing-Masing?

| **Situasi**                              | **Gunakan**            |
| ---------------------------------------- | ---------------------- |
| Data sering diubah                       | Pointer                |
| Data hanya dibaca                        | Struct langsung        |
| Struct kecil                             | Struct langsung        |
| Struct besar atau kompleks               | Pointer                |
| Performa sangat penting (heap vs. stack) | Pertimbangkan kasusnya |
| Banyak instance saling berbagi referensi | Pointer                |

---

## Ringkasan

Tidak semua struct perlu pointer karena **tidak semua kasus membutuhkan efisiensi memori atau modifikasi langsung**. Pointer adalah alat yang kuat, tapi menggunakannya di semua tempat akan membuat kode lebih sulit dibaca, meningkatkan risiko bug, dan menambah overhead pada runtime (misalnya garbage collection). Gunakan pointer dengan **tepat di tempat yang sesuai kebutuhan**.
