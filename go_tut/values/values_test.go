package values

import "testing"

func TestValues(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		integers()
	})

	t.Run("Floats", func(t *testing.T) {
		floats()
	})

	t.Run("Strings", func(t *testing.T) {
		strings()
	})

	t.Run("bool", func(t *testing.T) {
		bool()
	})
}
