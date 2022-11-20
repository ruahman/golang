package main

import (
	"C"
)

func main() {}

//export addOne
func addOne(val float64) float64 {
	return val + 1
}
