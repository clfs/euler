package p11

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 70600674
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
