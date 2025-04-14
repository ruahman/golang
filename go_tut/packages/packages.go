package packages

// packages is a group of files in the same folder

// we can import other packages
// import work per file
import (
	"fmt"
	"math"
	"math/rand"

	vars "ruahman.org/go_tut/variables"
)

// this is an exported function, a public Function
// this function can be called outside the package
func Add(x int, y int) int {
	return x + y
}

// this is a package function, a private function
// it can only be called inside the package
func packageFunc() {
	fmt.Println("package function")
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

var initMe string

// this gets run each time the pacakge is imported
func init() {
	initMe = "init function was ran"
}

func Packages() {
	vars.Variables()
	fmt.Println(initMe)
	packageFunc()
	fmt.Println("My favorite number is", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(Add(2, 3))
	fmt.Println(Add2(3, 4))
	fmt.Println(Swap("hello", "world"))
	fmt.Println(Swap2("hello", "world"))
}
