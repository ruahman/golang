package main

import (
	"go_tut/arrays"
	"go_tut/channels"
	"go_tut/custom_type"
	"go_tut/functions"
	"go_tut/interfaces"
	"go_tut/loops"
	"go_tut/maps"
	"go_tut/slices"
	"go_tut/structs"
	"go_tut/variables"
)

func main() {
	arrays.Demo()
	variables.Demo()
	functions.Demo()
	custom_type.Demo()
	slices.Demo()
	loops.Demo()
	structs.Demo()
	maps.Demo()
	interfaces.Demo()
	channels.Demo()
}
