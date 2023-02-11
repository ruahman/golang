package structs

import "fmt"

// go doesn't have inheretance but it does have composition
type contactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	age       int
	// if struct is embeded then we can access props directly
	contactInfo
}

// value receivers
// pass person by value, make no changes to struct
func (p Person) print() {
	fmt.Printf("%+v", p)
}

// pointer receivers
// pass by ref, can change struct
func (ptr *Person) updateName(newFirstName string) {
	(*ptr).firstName = newFirstName
}

// return a pointer
func newPerson(name string) *Person {
	p := Person{firstName: name}
	p.age = 42
	return &p
}

// inheritence from person
type Employee struct {
	Person
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

	var alex2 Person
	alex2.firstName = "andy"
	alex2.lastName = "vila"
	fmt.Printf("%+v", alex2)

	jim := Person{
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
