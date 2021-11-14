package p2

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 4613732
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
