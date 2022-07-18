package arrays

import "fmt"

func Demo() {
	// array has a fixed length,
	// you specify size in creation.
	var a [5]int

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a, b)
}
