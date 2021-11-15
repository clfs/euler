package p14

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 837799
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
