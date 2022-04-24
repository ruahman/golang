package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := person{
		firstName: "diego",
		lastName:  "vila",
		age:       40,
	}

	fmt.Println(p)

	type animal struct {
		name            string
		characteristics []string
	}

	b := animal{
		name:            "dog",
		characteristics: []string{"fun", "cut", "funny"},
	}

	fmt.Println(b)

	// promotion
	type herbivore struct {
		animal
		eatHuman bool
	}

	herb := herbivore{
		animal: animal{
			name:            "cow",
			characteristics: []string{"fat", "loud", "smelly"},
		},
		eatHuman: false,
	}

	fmt.Println(herb)
	fmt.Println(herb.name)

	for k, v := range herb.characteristics {
		fmt.Println(k, v)
	}

}
