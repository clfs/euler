package p12

import "testing"

func TestNDivisors(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{3, 2},
		{6, 4},
		{10, 4},
		{15, 4},
		{21, 4},
		{28, 6},
		{49, 3},
	}
	for _, c := range cases {
		if got := NDivisors(c.in); got != c.want {
			t.Errorf("NDivisors(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 76576500
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
