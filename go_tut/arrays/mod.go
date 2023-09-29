package arrays

import "fmt"

func Demo() {
	fmt.Println("----- arrays -----")
	// array has a fixed length,
  // they are always given a specific length
	// you specify size in creation.
	var a [5]int
	var x = [3]int{1, 2, 3}

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a, b, x)

	// infered lengths
	var arr1 = [...]int{1, 2, 3} 
	fmt.Println(arr1)

  // builtin len returns the lenght of an array
	fmt.Println(len(arr1))
}
