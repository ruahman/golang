// this means that this is an executable package, rather than a package
package main

import (
	"fmt"
)

func init() {
	fmt.Println("hello from init")
}

func main() {
	fmt.Println("Hello, World!")
}
