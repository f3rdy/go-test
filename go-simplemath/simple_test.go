package simplemath

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		in_a, in_b, want int
	}{
		{1, 1, 2},
		{3, 3, 6},
		{0, 1, 1},
		{5, 5, 10},
	}
	for _, c := range cases {
		got := Add(c.in_a, c.in_b)
		if got != c.want {
			t.Errorf("Add(%d,%d) == %d, want %d", c.in_a, c.in_b, got, c.want)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		in_a, in_b, want int
	}{
		{2,1,1},
		{5,1,4},
		{7,5,2},
	}
	for _, c := range cases {
		got := Sub(c.in_a, c.in_b)
		if got != c.want {
			t.Errorf("Sub(%d,%d) == %d, want %d", c.in_a, c.in_b, got, c.want)
		}
	}
	
}
