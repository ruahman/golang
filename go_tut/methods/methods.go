package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// a method is a function with a special receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (t MyFloat) Abs() float64 {
	if t < 0 {
		return float64(-t)
	}
	return float64(t)
}

// Methods with pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Exec() {
	v := Vertex{3, 4}
	println(v.Abs())
	t := MyFloat(3.24)
	fmt.Println(t.Abs())
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())
}
