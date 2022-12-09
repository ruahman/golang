package variables

import (
	"testing"
)

func TestVariables(t *testing.T) {
	var res Variables = variables()

	if res.card != "Ace of Spades" {
		t.Errorf("card is %s", res.card)
	}

	if res.card2 != "Five of Diamonds" {
		t.Errorf("card2 is %s", res.card2)
	}

	if res.a != 1 || res.b != 3 || res.c != 5 || res.d != 7 {
		t.Errorf("there was a problem with either a, b, c, d")
	}
}
