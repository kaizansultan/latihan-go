# Interface

> Fungsi interface Go apa dalam proyek nyata?

Interface dalam Go sangat berguna dalam proyek nyata, terutama ketika Anda membangun aplikasi yang membutuhkan **abstraksi**, **fleksibilitas**, dan **pengujian** yang lebih mudah. Berikut adalah beberapa alasan mengapa interface penting dalam pengembangan perangkat lunak dan contoh penerapannya dalam proyek nyata:

## Abstraksi dan Pengurangan Ketergantungan

Interface memungkinkan Anda untuk mendefinisikan kontrak tanpa harus memikirkan implementasinya langsung. Ini membantu mengurangi ketergantungan langsung antara komponen-komponen dalam sistem.

### Contoh:

Misalnya, jika Anda memiliki sistem yang berinteraksi dengan berbagai jenis penyimpanan (database, file system, atau cloud storage), Anda dapat mendefinisikan interface untuk penyimpanan yang menyederhanakan interaksi dengan penyimpanan tersebut.

```go
type Storage interface {
    Save(data string) error
    Load() (string, error)
}

type FileStorage struct {
    // implementasi khusus untuk file system
}

func (f *FileStorage) Save(data string) error {
    // menyimpan data ke file
    return nil
}

type DBStorage struct {
    // implementasi khusus untuk database
}

func (d *DBStorage) Save(data string) error {
    // menyimpan data ke database
    return nil
}
```

Dalam hal ini, Anda tidak perlu memikirkan jenis penyimpanan yang digunakan, selama tipe penyimpanan tersebut mengimplementasikan interface `Storage`. Ini mengurangi ketergantungan langsung pada tipe implementasi tertentu dan mempermudah perubahan atau penambahan sistem penyimpanan baru tanpa merusak kode lainnya.

## Fleksibilitas dalam Penggunaan

Interface memberi Anda fleksibilitas untuk mengganti implementasi tanpa mengubah kode yang bergantung pada interface tersebut. Anda dapat mengganti implementasi dengan tipe baru yang memiliki metode yang sesuai dengan interface, tanpa merusak fungsionalitas sistem.

### Contoh:

Misalnya, Anda memiliki sebuah aplikasi yang mengirimkan email. Anda bisa menggunakan interface untuk mendefinisikan pengiriman email tanpa bergantung pada implementasi spesifik (misalnya, menggunakan SMTP atau layanan pihak ketiga).

```go
type EmailSender interface {
    SendEmail(to, subject, body string) error
}

type SMTPService struct {
    // implementasi pengiriman email via SMTP
}

func (s *SMTPService) SendEmail(to, subject, body string) error {
    // kirim email via SMTP
    return nil
}

type MockEmailService struct {
    // implementasi untuk pengujian, misalnya hanya mencetak ke konsol
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
    fmt.Println("Email sent to:", to)
    return nil
}
```

Dalam contoh ini, Anda dapat mengganti `SMTPService` dengan `MockEmailService` untuk pengujian tanpa mengubah kode lainnya.

## Pengujian dan Mocking

Interface sangat berguna dalam unit testing. Anda dapat menggunakan interface untuk menggantikan bagian dari kode yang bergantung pada sumber eksternal, seperti database, API, atau layanan eksternal, dengan implementasi "mock" yang lebih mudah dikendalikan.

### Contoh:

Jika Anda ingin menguji bagian kode yang berinteraksi dengan database, Anda bisa membuat implementasi mock dari interface `Storage` dan menggunakannya untuk pengujian.

```go
type MockStorage struct {
    data string
}

func (m *MockStorage) Save(data string) error {
    m.data = data
    return nil
}

func (m *MockStorage) Load() (string, error) {
    return m.data, nil
}
```

Dengan cara ini, Anda bisa menguji logika aplikasi tanpa benar-benar berinteraksi dengan database. Hal ini mempercepat proses pengujian dan membuat pengujian lebih terisolasi dan dapat diulang.

## Dependency Injection

Interface memungkinkan **dependency injection**, yaitu pola di mana ketergantungan (dependency) dari suatu objek disuntikkan ke dalam objek tersebut dari luar, alih-alih objek tersebut yang membuat ketergantungannya sendiri. Ini sangat membantu dalam membuat kode lebih modular dan mudah diuji.

### Contoh:

Jika Anda memiliki aplikasi web yang bergantung pada berbagai layanan (seperti logging, otentikasi, dan penyimpanan data), Anda dapat mendefinisikan interface untuk layanan-layanan tersebut dan menyuntikkannya ke dalam komponen lain.

```go
type Logger interface {
    Log(message string)
}

type App struct {
    logger Logger
}

func NewApp(logger Logger) *App {
    return &App{logger: logger}
}

func (a *App) Run() {
    a.logger.Log("App is running!")
}
```

Dengan dependency injection, Anda dapat memberikan implementasi yang berbeda untuk `Logger` (misalnya, `FileLogger`, `ConsoleLogger`, atau `MockLogger` untuk pengujian).

## Polimorfisme

Interface memungkinkan polimorfisme dalam Go, yaitu kemampuan untuk menggunakan tipe yang berbeda secara konsisten jika mereka mengimplementasikan interface yang sama. Ini meningkatkan fleksibilitas dan kemampuan untuk menulis kode yang lebih generik.

### Contoh:

Anda bisa menggunakan interface untuk mendefinisikan berbagai tipe objek yang dapat diproses oleh fungsi yang sama.

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
```

Dengan `Shape` interface, Anda bisa menggunakan tipe `Circle` dan `Rectangle` dalam fungsi yang sama (`printArea`), meskipun keduanya adalah tipe yang berbeda.

## Mengurangi Redundansi

Dengan menggunakan interface, Anda bisa menghindari pengulangan kode (redundansi) karena interface memungkinkan Anda untuk menulis kode yang lebih generik yang bisa menangani berbagai tipe yang berbeda.

---

## Kesimpulan

Interface di Go berguna dalam proyek nyata karena mereka menyediakan cara untuk mendefinisikan **kontrak** dan **abstraksi**. Dengan menggunakan interface, Anda dapat:

-   Mengurangi ketergantungan langsung antar komponen,
-   Memungkinkan penggantian implementasi dengan mudah,
-   Membuat kode lebih fleksibel dan lebih mudah diuji,
-   Menghindari pengulangan kode dan meningkatkan modularitas.

Ini sangat penting dalam pengembangan perangkat lunak skala besar dan kompleks, di mana Anda perlu menjaga kode tetap bersih, fleksibel, dan mudah dipelihara.
