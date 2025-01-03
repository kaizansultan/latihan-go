# Exported Names

Di dalam bahasa Go, konsep **exported** dan **unexported** merujuk pada apakah suatu nama (seperti variabel, fungsi, atau tipe) dapat diakses dari luar paket tempat nama tersebut didefinisikan.

## Exported Name

-   Sebuah nama dianggap **"dieksport"** (exported) jika dimulai dengan huruf kapital.
-   Misalnya, jika kamu mendeklarasikan variabel atau fungsi dengan nama `Pizza`, `Pi`, `Area`, atau tipe data `Circle`, semua ini bisa diakses dari luar paket.
    -   Contoh: `Pi` yang ada dalam paket `math` adalah exported, sehingga bisa digunakan di luar paket `math`.
    -   Contoh kode:
        ```go
        import "math"
        fmt.Println(math.Pi) // Pi adalah exported, bisa diakses
        ```

## Unexported Name

-   Sebuah nama dianggap **"tidak dieksport"** (unexported) jika dimulai dengan huruf kecil.
-   Nama-nama yang tidak dieksport hanya dapat diakses dalam paket tempat mereka didefinisikan, dan tidak bisa diakses dari paket lain.

    -   Misalnya, `pizza` adalah variabel yang tidak dieksport, sehingga tidak bisa diakses di luar paket tempat `pizza` dideklarasikan.
    -   Contoh kode:

        ```go
        package main

        var pizza = "delicious"

        func main() {
            fmt.Println(pizza) // Tidak ada masalah karena pizza ada dalam paket yang sama
        }
        ```

## Contoh Kasus:

Dalam instruksi yang Anda berikan, ada contoh seperti ini:

-   `math.pi` adalah nama yang **tidak dieksport** (unexported) karena dimulai dengan huruf kecil. Maka, ketika Anda mencoba mengakses `math.pi` di luar paket `math`, akan muncul error.
-   `math.Pi` adalah nama yang **dieksport** (exported) karena dimulai dengan huruf kapital. Jadi, Anda bisa mengaksesnya di luar paket `math` tanpa masalah.

## Menyelesaikan Error:

Jika Anda mencoba mengakses nama `pi` yang tidak dieksport, Go akan memberikan error, seperti berikut:

```go
fmt.Println(math.pi) // Error: math.pi is not exported
```

Untuk memperbaiki ini, Anda harus mengganti `math.pi` menjadi `math.Pi`, seperti berikut:

```go
fmt.Println(math.Pi) // Tidak ada error: Pi adalah exported
```

## Kesimpulan:

-   **Exported** berarti dapat diakses dari luar paket, dan nama variabel atau fungsi harus dimulai dengan huruf kapital.
-   **Unexported** berarti hanya bisa diakses dalam paket itu sendiri, dan nama variabel atau fungsi dimulai dengan huruf kecil.
