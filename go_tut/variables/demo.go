package variables

import "fmt"

var noPublica string // private
var Publica string   // public

func Demo() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	noPublica = "not public"
  Publica = "public"

	fmt.Println(card)
	fmt.Println(noPublica, Publica)
	fmt.Println("hello from variables package")
}
