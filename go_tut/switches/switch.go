package switches

import "fmt"

func Demo() {
	fmt.Println("----- switches -----")

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("not a weekday")
	}

	switch day {
	case 1, 2, 3:
		fmt.Println("stuff")
	case 4, 5, 6:
		fmt.Println("other")
	}

	// type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I don't know what I am", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
