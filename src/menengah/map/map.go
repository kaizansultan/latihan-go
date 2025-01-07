package main

import "fmt"

type Orang struct {
	alamat  string
	panjang uint8
}

var siswa map[string]Orang

func main() {

	siswa = make(map[string]Orang)
	siswa["Vanesa Putri"] = Orang{
		"Uranus", 0,
	}

	fmt.Println(siswa["Vanesa Putri"])

	// Map literal, seperti struct literal, hanya sanya memerlukan kunci.
	guru := map[string]Orang{
		"Budi Utama": {
			"Halmahera", 16,
		},
		"Sri Utama": {
			"Cebu", 0,
		},
	}
	fmt.Println(guru)

	siswa["Vanesa Putri"] = Orang{"Jupiter", 0}
	fmt.Println(siswa)

	delete(guru, "Budi Utama")
	fmt.Println(guru)

	elem, ok := guru["Sri Utama"]
	fmt.Println(elem, ok)

	elem, ok = guru["Budi Utama"]
	fmt.Println(elem, ok)
}
