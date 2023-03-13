package slices

import "fmt"

func Demo() {
	fmt.Println("------------ slices demo -------------------")
	// Arrays are fixed, while slices can grow
	// when ever we create a slice an array is created
	// in the background
	// slices can grow and shrink as you see fit.
  // Giving a more powerful interface to sequences than arrays.


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

	// slices create an array in the background
	// both of these slices are pointing to the same array
	testSlice := []int{1, 2, 3, 4, 5}
	testSlice2 := testSlice[:3]
	testSlice2[0] = 666

	fmt.Println(testSlice, testSlice2)

	// lenght, capacity
	fmt.Println(len(testSlice2), cap(testSlice2))

  s := make([]string, 3)
	fmt.Println(s, len(s))

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  fmt.Println("len:", len(s))

  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  l := s[2:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)
  

}
