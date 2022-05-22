package main

import "fmt"

func main() {
  
  test := struct {
    firstName string
    lastName string
    age int
  }{
    firstName: "diego",
    lastName: "vila",
    age: 40,
  }

  fmt.Println(test) 
}

