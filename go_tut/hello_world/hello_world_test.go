package hello_world

import "testing"

// go test -v ./hello_world
func TestHelloWorld(t *testing.T) {
	t.Run("say hello world", func(t *testing.T) {
		HelloWorld()
	})

	// these are subtests
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Diego", "")
		want := "Hello, Diego"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Diego", "Spanish")
		want := "Hola, Diego"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Diego", "French")
		want := "Bonjour, Diego"
		assertCorrectMessage(t, got, want)
	})
}

// my custom assert function / helper function
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper
	// this will help other developers track down problems easier
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
