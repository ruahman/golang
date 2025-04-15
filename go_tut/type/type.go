package types

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// type alias
type (
	integer        = int
	sliceOfStrings = []string
)

var x integer

// type defenition
// you can add methods to you type
type (
	json     map[string]string
	distance float64
)

// you can add methods to types
func (miles distance) ToKm() distance {
	return distance(1.69 * miles)
}

func test() {
	// distance is an actual type not an alias
	d := distance(4.5)
	k := d.ToKm()
	fmt.Println(d, k)
}

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// bs = byte slice, err = error
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// receiver function???
// by convention, the receiver is a one or two letter abbreviation of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Types() {
	test()
	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Cards:")
	remainingCards.print()

	fmt.Println("Deck as string:")
	fmt.Println(cards.toString())

	fmt.Println("Saving deck to file...")
	_ = cards.saveToFile("my_cards.txt")

	fmt.Println("Reading deck from file...")
	cardsFromFile := newDeckFromFile("my_cards.txt")
	cardsFromFile.print()

	cards.shuffle()
	cards.print()
}
