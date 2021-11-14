package p7

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := int64(104743)
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
