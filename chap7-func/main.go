package main

import (
  "fmt"
)

func main() {
  // call this at the end
  defer LastHi()

  // push on defer stack
  defer func(){
    fmt.Println("Almost last??")
  }()


  a := HelloStr

  fmt.Println(a())

  b := TwoValues

  x, y := b()

  fmt.Println(x,y)
}


func HelloStr() string {
  return "Hello String"
}

func TwoValues() (string, string) {
  return "string2", "string2"
}

func LastHi() {
  fmt.Println("Last Hi")
}
