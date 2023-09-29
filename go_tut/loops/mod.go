package loops

import "fmt"

func Run() {
	fmt.Println("---- loops demo -----")
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	n := 1
	for n < 5 {
		n++
	}
	fmt.Println(n)

	// infinit loop???
	test := 1
	for {
		test++
		if test > 10 {
			break
		}
		fmt.Println(test)
	}

	// range
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println(idx, val)
	}
}
