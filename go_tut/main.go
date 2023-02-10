// run: go mod init <name-of-your-project>
// ex: go mod init go_tut
// - creates go.mod file with name of project and version of go

// this is the main package that runs on the terminal
// - in go overything is organized into packages
package main

// you have to explictily tell that the pacakage in inside the project package
import (
	"go_tut/arrays"
	// "go_tut/channels"
	"go_tut/conditions"
	// "go_tut/custom_type"
	// "go_tut/err"
	"go_tut/functions"
	// "go_tut/generics"
	"go_tut/goroutines"
	"go_tut/interfaces"
	"go_tut/loops"
	"go_tut/maps"
	"go_tut/pointers"
	// "go_tut/ranges"
	"go_tut/channels"
	"go_tut/slices"
	"go_tut/strings"
	"go_tut/structs"
	"go_tut/switches"
	// "go_tut/err"
	"go_tut/go_by_example"
	"go_tut/io"
	"go_tut/runes"
	"go_tut/variables"
)

// this is the main entry point for main package
func main() {

	io.Demo()
	variables.Demo()
	conditions.Demo()
	arrays.Demo()
	slices.Demo()
	loops.Demo()
	// custom_type.Demo()
	switches.Demo()
	strings.Demo()
	// ranges.Demo()
	functions.Demo()
	structs.Demo()
	interfaces.Demo()
	maps.Demo()
	pointers.Demo()
	// generics.Demo()
	// err.Demo()
	goroutines.Demo()
	channels.Demo()
	runes.Demo()

	go_by_example.HelloWorld()
}
