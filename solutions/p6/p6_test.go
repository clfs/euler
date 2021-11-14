package p6

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 25164150
	if got := Solve(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
