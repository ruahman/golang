package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeate := Repeat("a")
	expected := "aaaaa"

	if repeate != expected {
		t.Errorf("expected %q but got %q", expected, repeate)
	}
}

// to run the benchmark test, run the following command:
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
