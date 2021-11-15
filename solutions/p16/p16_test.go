package p16

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 1366
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
