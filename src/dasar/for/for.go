package main

import "fmt"

func main() {
	sum := 0

	// Go hanya memiliki satu loop keyword, yaitu `for`.

	// Terdiri dari init; conditional; post statement.
	for i := 0; i < 10; i++ {
		sum += 1
		println(sum)
	}

	// init dan post statement opsional.
	// Dengan begini, `for` adalah `while`.
	for sum < 200 {
		sum += sum
		println(sum)
	}

	// Jika Anda membuang loop condition, maka loop tidak akan berhenti.

	fmt.Println(sum)
}
