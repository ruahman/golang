package main

import (
  "fmt"
)

type Emplyee struct {
  FirstName string
  LastName string
}


func fullName(first string, last string) (fullName string) {
  fullName = fmt.Sprintf("%s %s",first,last)
  return fullName
}

func (e Emplyee) fullName() {
  fmt.Println(e.FirstName + " " + e.LastName)
}

func (e *Emplyee) changeLastName(lastName string) {
  e.LastName = lastName
}

func main() {

  e := Emplyee{
    FirstName: "Diego",
    LastName: "Vila",
  }

  r := fullName(e.FirstName, e.LastName)  
  fmt.Println(r)

  e.fullName()

  e.changeLastName("Vila Bennett")

  e.fullName()

  // recievers is a a better way
   



}
