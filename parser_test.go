package main

import "testing"

func TestParseExpression(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"512", 512},
		{"12 * 12", 144},
		{"10 * 120", 1200},
		{"5 + 5", 10},
		{"125 + 1", 126},
		{"125 * 125", 15625},
		{"5 - 5", 0},
		{"5 * 5", 25},
		{"5 / 5", 1},
		{"2 + 3 * 4 + 5", 19},
	}

	for _, c := range cases {
		p := newParser(c.in)
		got := p.parseExpression()
		if got != c.want {
			t.Errorf("%q != %d, want %d", c.in, got, c.want)
		}
	}
}
