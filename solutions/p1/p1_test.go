package p1

import "testing"

func TestP1(t *testing.T) {
	t.Parallel()
	want := 233168
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
