package p10

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := int64(142913828922)
	if got := Solve(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
