package values

import "fmt"

func strings() {
	fmt.Println("go" + "lang")
}

func integers() {
	fmt.Println("1 + 1 = ", 1+1)
}

func floats() {
	fmt.Println("7.0/3.0 =", 7.0/3.0)
}

func bool() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
