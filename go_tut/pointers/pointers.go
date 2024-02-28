package pointers

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Run() {
	fmt.Println("----- pointer demo -----")

	var x int = 3
	var px *int = &x
	fmt.Println(x, px, *px)

	var py *int = new(int)
	*py = 3
	fmt.Println(py, *py)

	i := 1
	fmt.Println(i)

	// doen't change i
	zeroval(i)
	fmt.Println(i)

	// changes i
	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)

	a := 5
	b := &a
	fmt.Println(a, b, *b)
	*b = 10
	fmt.Println(*b)

	// a pointer holds the memory address of a value
	var p *int
	c := 42
	p = &c
	fmt.Println(p, *p)
}
