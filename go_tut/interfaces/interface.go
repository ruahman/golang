package interfaces

import (
	"fmt"
	"math"
	"strings"
)

// an interface is a list fo methods
type Course struct {
	Title string
}

type Workshop struct {
	Name string
	Course
}

func (c Course) SignUp() bool {
	return true
}

func (w Workshop) SignUp() bool {
	return true
}

// an interface is just a list of methods
// this is implicit
type SignUpAble interface {
	SignUp() bool
}

// interfaces are data structures that are used to represent
// a set of methods for a type of struct

// all the types that have the following methods defined for it,
// fit this interface,
// for example if a type has the getGreating() method defined for it,
// callifes for the bot interface

// this is an interface type,
// we can't create a value of this type.
// we can only asign a value to this type that implements the methods defined for it.
// interfaces are implicit, we don't have to explicitly say that a type implements an interface.
// they just check if the methods are defined for the type.
// if the type matches the methods defined for the interface, then it is
// implicitly considered to implement the interface.
// you don't need to explicitly say that a type implements an interface.
type bot interface {
	getGreating() string
}

// this is a value type,
// we can create a value of this type.
type englishBot struct{}

type spanishBot struct{}

// use the interface bot as parameter,
// so that we can use any type that implements the bot interface
func printGreating(b bot) {
	fmt.Println(b.getGreating())
}

// receiver function for englishBot
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

func Interface() {
	fmt.Println("----- interfaces -----")

	var array_of_interfaces [2]SignUpAble
	array_of_interfaces[0] = Workshop{Name: "my workshop"}
	array_of_interfaces[1] = Course{Title: "my course"}

	for _, item := range array_of_interfaces {
		fmt.Println(item)
	}

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
