package main

import (
	"go_tut/arrays"
	"go_tut/channels"
	"go_tut/conditions"
	"go_tut/custom_type"
	"go_tut/functions"
	"go_tut/interfaces"
	"go_tut/loops"
	"go_tut/maps"
	"go_tut/slices"
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
	functions.Demo()
	custom_type.Demo()
	structs.Demo()
	maps.Demo()
	interfaces.Demo()
	channels.Demo()
}
