package maps

import "fmt"

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}

func Demo() {
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

	fmt.Println("----- c ------")
	fmt.Println(colors, colors3)

	printMap(colors)
}
