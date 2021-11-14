package p9

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 31875000
	if got := Solve(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
