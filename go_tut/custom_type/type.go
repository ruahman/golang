package custom_type

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func deckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
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

func Demo() {
  fmt.Println("***** custome type demo *****")
	demo := deck{"one", "two", "three"}
	fmt.Println("test")
	fmt.Println(demo)
	fmt.Println("tostring join: ", demo.toString())
	demo.saveToFile("test.txt")
	newDeck := deckFromFile("test.txt")
	newDeck.shuffle()
  fmt.Print(newDeck)
	demo.print()
}
