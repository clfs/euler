package p4

import "testing"

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 906609

	got, err := Solve()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
