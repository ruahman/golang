package dependency_injection

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("buffer", func(t *testing.T) {
		// buffer implements io.Writer
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("stdout", func(_ *testing.T) {
		// can't test stdout
		Greet(os.Stdout, "Chris")
	})
}
