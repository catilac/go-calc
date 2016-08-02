package main

import "testing"

func TestIsSpace(t *testing.T) {
	cases := TrueRunes(setSpaces)
	for _, c := range cases {
		got := isSpace(c.in)
		if got != c.want {
			t.Errorf("isSpace(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsOp(t *testing.T) {
	cases := TrueRunes(setOps)
	for _, c := range cases {
		got := isOp(c.in)
		if got != c.want {
			t.Errorf("isOp(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsNum(t *testing.T) {
	cases := TrueRunes(setNum)
	for _, c := range cases {
		got := isNum(c.in)
		if got != c.want {
			t.Errorf("isNum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPeek(t *testing.T) {
	l := newLexer("92 * 924")
	pos := l.pos
	got := l.peek()
	want := tokMul
	if got != tokMul {
		t.Errorf("peek() == %d, want %d", got, want)
	}

	if pos != l.pos {
		t.Errorf("l.pos == %d, want %d", l.pos, pos)
	}
}

func TestNext(t *testing.T) {
	input := "55 * 124"
	l := newLexer(input)
	for _, c := range input {
		got := l.next()
		if c != got {
			t.Errorf("expected %q, got %q", c, got)
		}
	}

	if l.pos != len(input) {
		t.Errorf("l.pos = %q, expected %q", l.pos, len(input))
	}

	got := l.next()
	if got != eof {
		t.Errorf("l.next() = %q, expected %q", got, eof)
	}
}

func TestRewind(t *testing.T) {
	l := newLexer("399 + 5")
	want := l.pos
	l.getTok()
	l.advance()
	l.rewind()

	if want != l.pos {
		t.Errorf("l.pos = %d, expected %d", l.pos, want)
	}
}

func TestGetTok(t *testing.T) {
	l := newLexer("512 * 123")

	for i := 0; i < 2; i++ {
		got, val := l.getTok()
		if got != tokNumber && val != "512" {
			t.Errorf("l.getTok() = %d, expected %d (iter: %d)", got, tokNumber, i)
		}
	}

	if l.pos != 0 {
		t.Errorf("l.pos = %d, expected %d", l.pos, 0)
	}

	if l.lastWidth != 4 {
		t.Errorf("l.lastWidth = %d, expected %d", l.lastWidth, 4)
	}
}

// Helpers
type runeSet int

const (
	setSpaces runeSet = iota
	setOps
	setNum
)

type runeCase struct {
	in   rune
	want bool
}

func TrueRunes(set runeSet) []runeCase {
	switch set {
	case setSpaces:
		return []runeCase{
			{' ', true},
			{'\t', true},
			{'\n', true},
			{'w', false},
		}
	case setOps:
		return []runeCase{
			{'+', true},
			{'-', true},
			{'*', true},
			{'/', true},
		}
	case setNum:
		return []runeCase{
			{'0', true},
			{'1', true},
			{'2', true},
			{'3', true},
			{'4', true},
			{'5', true},
			{'6', true},
			{'7', true},
			{'8', true},
			{'9', true},
		}
	default:
		panic("Invalid Rune Set")
	}
}
