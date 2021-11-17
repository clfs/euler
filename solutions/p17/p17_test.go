package p17

import "testing"

func TestWords(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   int
		want string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{10, "ten"},
		{11, "eleven"},
		{20, "twenty"},
		{99, "ninety-nine"},
		{101, "one hundred and one"},
		{115, "one hundred and fifteen"},
		{417, "four hundred and seventeen"},
		{342, "three hundred and forty-two"},
		{573, "five hundred and seventy-three"},
		{600, "six hundred"},
		{1000, "one thousand"},
	}
	for _, c := range cases {
		if got := Words(c.in); got != c.want {
			t.Errorf("Words(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSolve(t *testing.T) {
	t.Parallel()
	want := 21124
	if got := Solve(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
