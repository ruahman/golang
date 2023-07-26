package variables

import (
	"fmt"
	"reflect"
)

// private variable
var noPublica string = "noPublica"

// Public variable
var Publica string = "Publica"

const s string = "constant"

// if you want to declare a variable outside a function you need to use var
var (
	xx int = 1
	yy int = 3
	zz int = 66
)

func Exec() {
	fmt.Println(xx, yy, zz)
	fmt.Println(s)
	fmt.Println(noPublica)

	var student1 string = "John"
	student2 := "Jane"
	x := 2
	fmt.Println(student1, student2, x)

	student1 = "diego"
	fmt.Println(student1)

	// default values
	var a string
	var b int
	var c bool
	var d float64

	fmt.Printf("%v %v %v %v\n", a, b, c, d)

	// values assignment after declaration
	var student3 string // nolint
	student3 = "Diego"
	fmt.Println(student3)

	// multiple variable declaration
	var a1, b2, c3, d4 int = 1, 3, 5, 7
	fmt.Println(a1, b2, c3, d4)

	// different types in same declaration
	a5, b6 := 6, "hello"
	c7, d8 := 7, "world"
	fmt.Println(a5, b6, c7, d8)

	// variable delcaration block
	var (
		a9  int
		b10 int    = 1
		c11 string = "hello"
	)
	fmt.Println(a9, b10, c11)

	// show types
	fmt.Println(reflect.TypeOf(33))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello"))

	// casting
	cV1 := 3.16
	cV2 := int(cV1)
	fmt.Println(cV1, cV2)

	// show value and type
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has vale: %v and type: %T\n", i, i)
	fmt.Printf("j has vale: %v and type: %T\n", j, j)
}
