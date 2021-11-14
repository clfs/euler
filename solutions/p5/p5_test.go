package p5

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 232792560

	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
