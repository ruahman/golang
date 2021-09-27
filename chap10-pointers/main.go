package main

import (
  "fmt"
)

type Person struct {
  firstName string
  lastName string
}

func changeName(p *Person){
  p.firstName = "I changed"
}

func changeArray(a *[5]int){
  a[2] = 773
}

func changeArraySlice(a []int){
  a[3] = 878 
}

func main() {
  person := Person {
    firstName: "diego",
    lastName: "vila",
  }

  changeName(&person)
  fmt.Println(person)

  a := [5]int{}
  changeArray(&a)
  fmt.Println(a)

  as := []int{0, 2, 3, 4, 5}
  changeArraySlice(as)
  fmt.Println(as)

}
