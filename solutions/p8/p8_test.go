package p8

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 23514624000
	if got := Solve(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
