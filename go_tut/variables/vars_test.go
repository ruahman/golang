package variables

import (
	"testing"
)

// run "go test -v -run TestDemo ./variables"
func TestDemo(t *testing.T) {
	var card, card2, card3 = Demo()

	if card != "Ace of Spades" {
		t.Errorf("card is %s", card)
	}

	if card2 != "Ace of Spades2" {
		t.Errorf("card2 is %s", card2)
	}

	if card3 != "I was changed" {
		t.Errorf("card3 is %s", card3)
	}

	// if res.a != 1 || res.b != 3 || res.c != 5 || res.d != 7 {
	// 	t.Errorf("there was a problem with either a, b, c, d")
	// }
}
