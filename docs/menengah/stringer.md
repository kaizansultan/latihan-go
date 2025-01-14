# Stringers

Di Go, **Stringer** adalah sebuah interface yang didefinisikan dalam paket `fmt`. Interface ini digunakan untuk menentukan cara sebuah tipe data dikonversi menjadi string.

Definisinya sangat sederhana:

```go
type Stringer interface {
    String() string
}
```

Jika sebuah tipe data mengimplementasikan metode `String() string`, maka tipe tersebut secara otomatis diakui sebagai **Stringer**. Ketika tipe ini digunakan dengan fungsi dari `fmt` seperti `fmt.Print` atau `fmt.Println`, metode `String` akan dipanggil secara otomatis untuk mendapatkan representasi string dari nilai tersebut.

## Contoh Sederhana

Berikut adalah tipe data yang mengimplementasikan interface `Stringer`:

```go
package main

import "fmt"

// Tipe Struct
type Person struct {
    Name string
    Age  int
}

// Implementasi metode String
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    // fmt.Println akan otomatis memanggil metode String
    fmt.Println(p) // Output: Alice (30 years old)
}
```

**Penjelasan:**

-   Struct `Person` memiliki dua field: `Name` dan `Age`.
-   Metode `String()` diimplementasikan untuk memberikan representasi string khusus bagi `Person`.
-   Ketika `p` dipanggil di dalam `fmt.Println`, Go memanggil `p.String()` untuk mendapatkan string `"Alice (30 years old)"`.

## Kapan dan Mengapa Menggunakan Stringer?

1. **Meningkatkan Readability**  
   Dengan `Stringer`, Anda dapat menentukan cara data ditampilkan secara lebih deskriptif, terutama saat debugging atau logging.

    **Tanpa Stringer:**

    ```go
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p) // Output default: {Alice 30}
    ```

    **Dengan Stringer:**

    ```go
    fmt.Println(p) // Output: Alice (30 years old)
    ```

2. **Kustomisasi Output**  
   Anda bisa membuat representasi string sesuai kebutuhan, misalnya untuk menyembunyikan data sensitif atau mengubah format data.

## Contoh Implementasi Lain

### 1. Mengubah Representasi Data

Misalnya, Anda memiliki tipe data `Point` untuk koordinat 2D:

```go
package main

import "fmt"

type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
    p := Point{X: 3, Y: 4}
    fmt.Println(p) // Output: (3, 4)
}
```

### Menyembunyikan Data Sensitif

Misalnya, Anda ingin menyembunyikan detail nomor kartu kredit:

```go
package main

import "fmt"

type CreditCard struct {
    Number string
}

func (cc CreditCard) String() string {
    return fmt.Sprintf("**** **** **** %s", cc.Number[len(cc.Number)-4:])
}

func main() {
    card := CreditCard{Number: "1234567812345678"}
    fmt.Println(card) // Output: **** **** **** 5678
}
```

## Cara Kerja di Paket `fmt`

Ketika Anda menggunakan fungsi seperti `fmt.Print` atau `fmt.Println`, Go akan:

1. Mengecek apakah tipe data yang dicetak mengimplementasikan metode `String() string`.
2. Jika ya, metode tersebut akan dipanggil untuk mendapatkan representasi string.
3. Jika tidak, Go menggunakan representasi bawaan, seperti:
    - `{}` untuk struct.
    - Tipe data default (misalnya angka, boolean).

---

## **Kesimpulan**

-   **Stringer** adalah interface sederhana untuk memberikan representasi string khusus pada tipe data.
-   Banyak digunakan di Go, terutama dalam debugging, logging, dan output yang lebih deskriptif.
-   Dengan `Stringer`, Anda dapat menentukan bagaimana data kompleks ditampilkan dalam bentuk string, menggantikan representasi bawaan Go yang terkadang kurang informatif.
