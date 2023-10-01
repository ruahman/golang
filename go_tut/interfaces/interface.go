package interfaces

import (
	"fmt"
	"math"
	"strings"
)

// interfaces are data structures that are used to represent
// a set of methods for a type of struct

type bot interface {
	getGreating() string
}

type englishBot struct{}

type spanishBot struct{}

func printGreating(b bot) {
	fmt.Println(b.getGreating())
}

func (englishBot) getGreating() string {
	return "Hi"
}

func (spanishBot) getGreating() string {
	return "Hola"
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 2 * (s.side + s.side)
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("String: " + strings.ToUpper(i.(string)))
	case int:
		fmt.Println("Int: ", i.(int))
	default:
		fmt.Println("something else")
	}
}

func explain2(i interface{}) {
	switch t := i.(type) {
	case string:
		fmt.Println("String: " + strings.ToUpper(t))
	case int:
		fmt.Println("Integer: ", t)
	}
}

type Abser interface {
	Abs() float64
}

// both these type impment the Abser interface
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// empty interfaces
// interface{} is the empty interface, and it defines no methods.
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Demo() {
	fmt.Println("----- interfaces -----")

	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)

	// interface Shape
	var s Shape
	s = Rect{width: 5.0, height: 4.0}
	fmt.Println(s.Area(), s.Perimeter())

	s = Square{side: 7.0}
	fmt.Println(s.Area(), s.Perimeter())

	explain("hello world")
	explain(23)

	explain2("hello world2")
	explain2(27)

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = v // a Vertex implements Abser
	fmt.Println(a.Abs())

	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
}
