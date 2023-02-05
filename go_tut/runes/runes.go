package runes

import (
	"fmt"
	"unicode/utf8"
)

func Demo() {
	// in go charaters are called runes
	rStr := "abcdefg"
	fmt.Println("count the number of runes: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	aStr := "efghijk"
	rArr := []rune(aStr)
	for _, x := range rArr {
		fmt.Println(x)
	}
}
