package arrays

import "fmt"

func Run() {
	fmt.Println("----- arrays -----")
	// array has a fixed length,
	// they are always given a specific length
	// you specify size in creation.
	var a [5]int
	x := [3]int{1, 2, 3}

	b := [5]int{1, 2, 3, 4, 5}

	// ... is used to infer the length of the array
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a, b, c, x)
	// show the length of each array
	fmt.Println(len(a), len(b), len(c), len(x))

	// infered lengths
	arr1 := [...]int{1, 2, 3}
	fmt.Println(arr1)

	// builtin len returns the lenght of an array
	fmt.Println(len(arr1))

	mulArray := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(mulArray)
}
