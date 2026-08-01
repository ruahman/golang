// this means that this is an executable package, rather than a package
package main

import (
	"fmt"

	"ruahman.org/golang/go_tut/hello_world"
)

func init() {
	fmt.Println("hello from init")
}

func main() {
	hello_world.HelloWorld()
}
