package strings

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Run() {
	const s = "สวัสดี"

	// length just returns the length in bytes. string hold bytes
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
	}

	// to find the number of characters in a string, use utf8.RuneCountInString
	// number of runes in string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// each character in a string is a rune
	// some characters in UTF-8 are more than one byte,
	// that is why we use RuneCountInString
	// and not len to find the number of characters in a string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	p := fmt.Println
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
}
