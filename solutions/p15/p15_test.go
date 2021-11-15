package p15

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := int64(137846528820)
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
