# Slice

## Slices Tidak Menyimpan Data Sendiri

Slice adalah _view_ atau _referensi_ ke suatu bagian dari array yang mendasarinya. Slice tidak menyimpan data sendiri, melainkan hanya mendeskripsikan bagian mana dari array yang ingin diakses.

**Contoh:**

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // slice mengacu pada elemen arr[1], arr[2], dan arr[3]
fmt.Println(slice) // Output: [2 3 4]
```

Di sini, `slice` tidak menyimpan data baru, tetapi hanya mengakses sebagian dari array `arr`.

## Mengubah Elemen Slice Mengubah Array yang Mendasarinya

Karena slice adalah referensi ke array, jika kita mengubah elemen dalam slice, maka elemen tersebut akan mengubah nilai yang ada dalam array yang mendasarinya.

**Contoh:**

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // slice mengacu pada arr[1], arr[2], arr[3]

slice[0] = 99 // Mengubah elemen pertama di slice (arr[1]) menjadi 99

fmt.Println(arr)   // Output: [1 99 3 4 5] (array terpengaruh)
fmt.Println(slice) // Output: [99 3 4] (slice juga terpengaruh)
```

-   **Mengubah elemen di slice**: Perubahan langsung terjadi pada array yang mendasari slice tersebut, karena slice hanya memegang referensi ke bagian dari array itu.
-   **Hasilnya**: Array yang mendasari slice akan terlihat perubahan tersebut, begitu juga dengan slice lain yang berbagi array yang sama.

## Slice Lain yang Mengakses Array yang Sama

Jika ada slice lain yang juga merujuk ke array yang sama, perubahan yang dilakukan pada salah satu slice akan terlihat pada slice lainnya.

**Contoh:**

```go
arr := [5]int{1, 2, 3, 4, 5}
slice1 := arr[1:4]  // Mengacu ke arr[1], arr[2], arr[3]
slice2 := arr[2:5]  // Mengacu ke arr[2], arr[3], arr[4]

slice1[0] = 99  // Mengubah arr[1]
fmt.Println(arr) // Output: [1 99 3 4 5]
fmt.Println(slice1) // Output: [99 3 4]
fmt.Println(slice2) // Output: [3 4 5] (terlihat perubahan pada arr[2])
```

-   **slice1** mengubah elemen `arr[1]`, yang akan mempengaruhi **slice2** juga, karena keduanya berbagi array yang sama.

## Kesimpulan

-   **Slice** adalah referensi ke array dan bukan salinan independen.
-   Mengubah elemen di dalam slice akan mempengaruhi array yang mendasari dan slice lain yang merujuk ke array yang sama.
-   Dengan kata lain, **slice berfungsi seperti pointer** ke bagian dari array.
