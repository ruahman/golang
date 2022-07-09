package slices

import "fmt"

func Demo() {
	fmt.Println("slices demo...")
	// Arrays are fixed, while slices can grow
	cards := []string{"Ace of Diamonds", "Ace of Spades"}
	cards = append(cards, "Six of Spades")
	cards = append(cards, "one", "two", "three")
	fmt.Println(cards)

	for idx, card := range cards {
		fmt.Println(idx, card)
	}

	for _, card := range cards {
		fmt.Println(card)
	}

	fmt.Println("ranges: ", cards[2:5])
}
