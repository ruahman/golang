package conditions

import "fmt"

func Demo() {
  fmt.Println("----- conditions -----")
	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	time := 20
	if time < 18 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

	time2 := 22
	if time2 < 10 {
		fmt.Println("Good morning")
	} else if time2 < 20 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}
}
