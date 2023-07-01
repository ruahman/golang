package hello_world

import "testing"

func TestHelloWorld(t *testing.T) {
	HelloWorld()
	t.Log("just run `go test -v ./hello_world` or inside your package `go test`")
}
