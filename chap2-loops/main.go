package main

import "fmt"

func main() {
  sum := 0
  for i := 1; i < 5; i++ {
    sum += i
  }
  fmt.Println(sum)

  n := 1
  for n < 5 {
    n++
  }
  fmt.Println(n)

  // infinit loop???
  test := 1
  for {
    test++
    if test > 10 {
      break
    }
    fmt.Println(test)
  }
}
