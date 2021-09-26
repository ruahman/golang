package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

  fmt.Println("a: ",a)

  // declare an array
  var b []int

  b = []int{ 2, 4, 6, 8, 10 }

  fmt.Println("b: ",b)

  c := make([]int, 4)
  c = append(c, 60, 40)
  
  fmt.Println(c)
  fmt.Println(c[2:])

  // range
  for i, v := range c {
    fmt.Println(i,v)
  }


}
