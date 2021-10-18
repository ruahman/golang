package main

import (
	"fmt"
)

func main() {
	var s, sep string
	test := [3]string{"hello", "golang", "world"}

	for i := 1; i < len(test); i++ {
		s += sep + test[i]
		sep = " "
	}

	fmt.Println(s)
}
