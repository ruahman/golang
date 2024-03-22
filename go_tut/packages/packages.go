package packages

import (
	"fmt"
	"math"
	"math/rand"
)

func Add(x int, y int) int {
	return x + y
}

func Add2(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

// named return variables,
// variables are delared in return type
func Swap2(x, y string) (a, b string) {
	a = y
	b = x
	return
}

func Run() {
	fmt.Println("My favorite number is", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(Add(2, 3))
	fmt.Println(Add2(3, 4))
	fmt.Println(Swap("hello", "world"))
	fmt.Println(Swap2("hello", "world"))
}
