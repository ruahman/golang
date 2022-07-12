package interfaces

import "fmt"

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

var s Shape

func Demo() {
	fmt.Println("----- interfaces -----")

	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)

	s = Rect{width: 5.0, height: 4.0}
	fmt.Println(s.Area(), s.Perimeter())

	s = Square{side: 7.0}
	fmt.Println(s.Area(), s.Perimeter())
}
