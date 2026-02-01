package struct_embedding

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// this kind compose a struct
type container struct {
	base
	str string
}

func StructEmbeding() {
	fmt.Println("struct embedding")
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// we can access the base fields directly
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// we can also use the full path
	fmt.Println("also num:", co.base.num)

	// since container embeds base, the method of base also become the methods of a container
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// container now mathes interface
	var d describer = co
	fmt.Println("describer:", d.describe())
}
