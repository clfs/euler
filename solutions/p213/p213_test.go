package p213

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	t.Skip("unfinished")
	want := 0.0
	if got := Solve(); got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
