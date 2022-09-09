package main

import (
	"go_tut/arrays"
	"go_tut/channels"
	"go_tut/conditions"
	"go_tut/custom_type"
	"go_tut/err"
	"go_tut/functions"
	"go_tut/generics"
	"go_tut/interfaces"
	"go_tut/loops"
	"go_tut/maps"
	"go_tut/pointers"
	"go_tut/ranges"
	"go_tut/slices"
	"go_tut/strings"
	"go_tut/structs"
	"go_tut/switches"
	"go_tut/variables"
)

func main() {
	variables.Demo()
	arrays.Demo()
	slices.Demo()
	conditions.Demo()
	switches.Demo()
	loops.Demo()
	ranges.Demo()
	pointers.Demo()
	strings.Demo()
	custom_type.Demo()
	functions.Demo()
	structs.Demo()
	generics.Demo()
	err.Demo()
	maps.Demo()
	interfaces.Demo()
	channels.Demo()
}
