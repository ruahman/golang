package main

import "fmt"

func main() {
	cards := []string{"foo", "bar"}
	cards = append(cards, "six of spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
