package hello_world

import "testing"

// go test -v ./hello_world
func TestHelloWorld(t *testing.T) {
	Run()
	t.Log("just run `go test -v ./hello_world` or inside your package `go test`")
}
