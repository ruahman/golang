// this means that this is an executable package, rather than a package
package main

import (
	"fmt"

	"ruahman.org/golang/go_tut/conditions"
	"ruahman.org/golang/go_tut/hello_world"
	"ruahman.org/golang/go_tut/variables"
)

func init() {
	fmt.Println("hello from init")
}

func main() {
	fmt.Println("Hello, World!")

	variables.Variables()
	conditions.Conditions()
	hello_world.HelloWorld()
}
