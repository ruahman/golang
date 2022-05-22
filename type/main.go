package main

import (
  "fmt"
  "strings"
)

func explain(i interface{}) {
  switch i.(type) {
    case string:
      fmt.Println("String: " + strings.ToUpper(i.(string)) )
    case int:
      fmt.Println("Int: ", i.(int))
    default:
      fmt.Println("something else")
  }
}

func explain2(i interface{}) {
  switch t := i.(type) {
    case string:
      fmt.Println("String: " + strings.ToUpper(t))
    case int:
      fmt.Println("Integer: ", t)
  }
}

func main() {
  explain("hello world")
  explain(23)

  explain2("hello world2")
  explain2(27)
}
