package variables

import (
	"fmt"
	"reflect"
)

// private variable
var noPublica string

// Public variable
var Publica string // public

const s string = "constant"

type Variables struct {
	card  string
	card2 string
	card3 string
	a     int
	b     int
	c     int
	d     int
}

func Demo() (string, string, string) {
	var card string = "Ace of Spades"
	var card2 = "Ace of Spades2"
	card3 := "Five of Diamonds" // only for initializing

	card3 = "I was changed" // if you want to change the variable then just us =

	fmt.Println(card, card2, card3)

	// show types
	fmt.Println(reflect.TypeOf(33))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello"))

	// casting
	cV1 := 3.16
	cV2 := int(cV1)
	fmt.Println(cV1, cV2)

	return card, card2, card3

	// // multiple variable declaration
	// var a, b, c, d int = 1, 3, 5, 7
	// fmt.Println("variables: ", variables())

	// // variable declaration block
	// var (
	// 	a1 int
	// 	b2 int    = 1
	// 	c3 string = "hello"
	// )

	// fmt.Println(a1, b2, c3)

	// fmt.Println(noPublica, Publica)
	// fmt.Println("hello from variables package")
	// fmt.Println(s)

	// const A = 1
	// const B int = 2

	// fmt.Println(A, B)

	// var i string = "Hello"
	// var j int = 15

	// fmt.Printf("i has vale: %v and type: %T\n", i, i)
	// fmt.Printf("j has vale: %v and type: %T\n", j, j)

	// var abool bool = true
	// var bint int = 5
	// var buint uint = 500
	// var cfloat float32 = 3.14
	// var cfloat64 float64 = 11.1
	// var dstr string = "Hi!"

	// fmt.Println(abool, bint, buint, cfloat, cfloat64, dstr)
}
