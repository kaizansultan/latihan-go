package main

import (
	"fmt"
	"math/rand"
)

/*
Bisa ditulis seperti ini:

```
import "fmt"
import "math/rand"
```

Tapi lebih baik ditulis dalam bentuk *factored import*.
*/

func main() {
	fmt.Println("Halo Guys! Aku suka angka", rand.Intn(10))
}
