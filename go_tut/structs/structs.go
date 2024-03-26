package structs

import "fmt"

// struct is a collection of fields

// go doesn't have inheretance but it does have composition
type contactInfo struct {
	email   string
	zipCode int
}

// order of fields matters, if we want to declare a struct without passing field names.
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
// go is a pass by value language, so we need to use pointers to change values
// go will automatically convert the person to a pointer if it sees a pointer receiver
// pass by ref, can change struct
func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
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

// struct embedding
type base struct {
	num int
}

func (b base) describe() {
	fmt.Println("describe base:", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe()
}

func Demo() {
	fmt.Println("---- structs -----")

	// with fields
	alexWithFields := Person{
		firstName:   "alex",
		lastName:    "lopez",
		contactInfo: contactInfo{email: "de@y.com", zipCode: 50506},
	}
	fmt.Println(alexWithFields)

	// without fields
	alex := Person{"alex", "lopez", 33, contactInfo{"dego@yahoo.com", 60647}}
	fmt.Println(alex)

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

	// go will automatically convert jim to a pointer
	// jimPtr := &jim
	// jimPtr.updateName("jimmyyyyyyyyyyyy")

	jim.updateName("Jjjjjjjjjjjimmmmyyyyyyyyyyyyyyyyyy")
	jim.print()

	test := newPerson("andy")
	fmt.Println(test)

	// struct pointer
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // this is fine even though p is a pointer
	fmt.Println(v)

	// struct embedding
	co := container{base: base{num: 666}, str: "foobar"}
	// you can access the base struct indirectly
	fmt.Println("co.num:", co.num, "co.str:", co.str)
	// or directly
	fmt.Println("co.base.num:", co.base.num)

	// you can call methods on the embedded struct
	co.describe()

	// can also calify for interfaces ment for the embedded struct
	var d describer = co
	d.describe()
}
