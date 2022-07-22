package slices

import "fmt"

func Demo() {
	fmt.Println("slices demo...")
	// Arrays are fixed, while slices can grow
	// when ever we create a slice an array is created
	// in the background
	// slices can grow and shrink as you see fit.

	cards := []string{"Ace of Diamonds", "Ace of Spades"}

	myslice := []int{1, 2, 3}
	fmt.Println(myslice)

	// you can append slices
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

	s := make([]string, 3)
	fmt.Println(s, len(s))

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslicex := arr1[2:4]
	fmt.Println(myslicex)

	// make slice with make
	myslice2 := make([]int, 5)
	fmt.Println(myslice2)

	// append to slices
	myslices3 := []int{1, 2, 3, 4, 5, 6}
	myslices3 = append(myslices3, 20, 21)
	fmt.Println(myslices3)

	myslice4 := []int{1, 2, 3}
	myslice5 := []int{4, 5, 6}
	myslice7 := append(myslice4, myslice5...)
	fmt.Println(myslice7)
}
