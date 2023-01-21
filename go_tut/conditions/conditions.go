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

	// asighn variable in if

	if name, age := "alex", 26; name == "alex" {
		fmt.Print(name, age)
	} else {
		fmt.Println("no name")
	}

	users := make(map[string]string)

	users["diego"] = "dego_vila@yahoo.com"
	users["andy"] = "abvila@gmail.com"

	if email, ok := users["diego"]; ok == true {
		fmt.Println("diego found", email)
	} else {
		fmt.Println("nothing")
	}
}
