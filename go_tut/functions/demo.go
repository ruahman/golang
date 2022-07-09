package functions

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func multiple_return() (int, bool) {
	return 123, false
}

func Demo() {
	card := newCard()
	i, b := multiple_return()
  fmt.Println("multiple retrun: ", i, b)
	fmt.Println(card)
	fmt.Println("demo from functions")
}
