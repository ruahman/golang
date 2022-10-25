package pointers

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Demo() {
  fmt.Println("----- pointer demo -----")
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)
}
