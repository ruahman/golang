package main

import "fmt"

func main() {

	var b string

	a := "hello"

	if a == "hello" {
		fmt.Println(a)
	} else {
		fmt.Println("this is the else")
	}

	if b == "" {
		fmt.Println("this is the default value")
	}
}
