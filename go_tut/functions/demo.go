package functions

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func multiple_return() (int, bool) {
	return 123, false
}

func lastHi() {
	fmt.Println("lastHi")
}

func Demo() {
	card := newCard()
	i, b := multiple_return()

	// call this at the end
	defer lastHi()

	// push on defer stack
	defer func() {
		fmt.Println("Almost last??")
	}()

	fmt.Println("multiple retrun: ", i, b)
	fmt.Println(card)
	fmt.Println("demo from functions")
}
