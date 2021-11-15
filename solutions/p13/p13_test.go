package p13

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 5537376230
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
