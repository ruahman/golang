package main

import "fmt"

func main() {
	fmt.Println("fmt.Run")

	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	greeting := fmt.Sprintf("Your name is %v and you are %v years old.", name, age)
	fmt.Println(greeting)
	fmt.Printf("type: %T\n", greeting)
}
