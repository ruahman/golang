package maps

// maps are just key maps of values

import "fmt"

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}

func Demo() {
	fmt.Println("***** map demo *****")

	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#FA0234",
		"purple": "#AF3GH",
	}

	// var colors2 map[string]string

	colors3 := make(map[string]string)
	colors3["yello"] = "hi"
	colors3["test"] = "remove me"

	delete(colors3, "test")

	fmt.Println(colors, colors3)

	printMap(colors)

	students := make(map[string][]int)

	students["Diego"] = []int{13, 45, 67}
	students["Alex"] = []int{34, 55, 66}

	// tho optional second retrun value can be used to disamiguate between
	// misssing keys and keys with zero values like 0 or ""
	_, found := students["andy"]
	fmt.Println("found: ", found)

	fmt.Println(students)
	fmt.Println(len(students))

	for k, v := range students {
		fmt.Println(k, v)
	}

	type Vertex struct {
		Lat, Long float64
	}

	m := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
	if x, ok := m["Google"]; ok {
		fmt.Println(x)
	}
	if y, ok := m["Bell Labs"]; ok {
		fmt.Println(y)
	}
	if z, ok := m["Apple"]; ok {
		fmt.Println(z)
	} else {
		fmt.Println("Apple not found")
	}
}
