package hello_world

import "testing"

// go test -v ./hello_world
func TestHello(t *testing.T) {
	// these are subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Diego", "")
		want := "Hello, Diego"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Diego", "Spanish")
		want := "Hola, Diego"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Diego", "French")
		want := "Bonjour, Diego"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// for helper functions it is a good idea to accept testing.TB interface
	// which is an interface that *testing.T and *testing.B both satisfy
	// so that your helper function can be used in either tests or benchmarks

	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper
	// this will help other developers track down problems easier
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
