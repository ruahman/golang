package constanats

import "fmt"

func constants() {
	// sort of like enums
	const (
		dimandCard = iota
		spadeCard
		heartCard
	)

	// skip zero
	const (
		_ = iota
		redKind
		blueKind
		yellowKind
	)

	// bytes
	const (
		_         = iota
		adminRole = 1 << iota
		publisherRole
		consumerRole
	)

	var role byte = adminRole | publisherRole
	fmt.Println(role)
	fmt.Println("does it have an admin role assigned to it? ", (adminRole&role) == adminRole)

	fmt.Println(dimandCard)
	fmt.Println(spadeCard)
	fmt.Println(heartCard)
}
