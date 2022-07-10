package structs

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (ptr *person) updateName(newFirstName string) {
	(*ptr).firstName = newFirstName
}

func Demo() {
	fmt.Println("---- hello structs")

	// alex := person{"alex", "lopez"}
	// diego := person{firstName: "diego", lastName: "vila"}

	// fmt.Println(alex, diego)

	var alex2 person
	alex2.firstName = "andy"
	alex2.lastName = "vila"
	fmt.Printf("%+v", alex2)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 60647,
		},
	}

	jim.print()

	// jimPtr := &jim
	// jimPtr.updateName("jimmyyyyyyyyyyyy")
  jim.updateName("Jjjjjjjjjjjimmmmyyyyyyyyyyyyyyyyyy")
	jim.print()
}
