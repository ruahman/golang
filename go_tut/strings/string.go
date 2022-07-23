package strings

import (
	"fmt"
	"unicode/utf8"
)

func Demo() {
	const s = "สวัสดี"

	// string hold bytes
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
	}
	// number of runes in string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runValue := range s {
		fmt.Printf("%#U starts at %d\n", runValue, idx)
	}

}
