package p3

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 6857
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
