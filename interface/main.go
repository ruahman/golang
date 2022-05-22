package main

import ("fmt")

type Shape interface {
  Area() float64
  Perimeter() float64
}

type Rect struct {
  width float64
  height float64
}

func (r Rect) Area() float64 {
  return r.width * r.height
}

func (r Rect) Perimeter() float64 {
  return 2 * (r.width + r.height)
}


var s Shape

func main() {

  s = Rect{width:5.0,height:4.0,}

  fmt.Println(s.Area(),s.Perimeter())

}

