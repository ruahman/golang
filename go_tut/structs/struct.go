package structs

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	// if struct is embeded then we can access props directly
	contactInfo
}

// pass person by value, make no changes to struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

// pass by ref, can change struct
func (ptr *person) updateName(newFirstName string) {
	(*ptr).firstName = newFirstName
}

// return a pointer
func newPerson(name string) *person {
	p := person{firstName: name}
	p.age = 42
	return &p
}

// inheritence from person
type Employee struct {
	person
}

func emp_func() {
	e := Employee{}
	e.firstName = "diego"
	e.lastName = "vila"
	e.age = 42
	fmt.Println(e)
}

func Demo() {
	fmt.Println("---- structs -----")

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
	// this is okay because struct is embeded
	fmt.Println(
		jim.contactInfo.email,
		jim.email,
		jim.contactInfo.zipCode,
		jim.zipCode)

	// jimPtr := &jim
	// jimPtr.updateName("jimmyyyyyyyyyyyy")
	jim.updateName("Jjjjjjjjjjjimmmmyyyyyyyyyyyyyyyyyy")
	jim.print()

	test := newPerson("andy")
	fmt.Println(test)
}
