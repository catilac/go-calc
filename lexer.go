package main

import (
	"unicode/utf8"
)

type Token int

const eof = -1
const (
	tokError Token = iota
	tokEOF

	tokNumber

	// + - / *
	tokAdd
	tokSub
	tokDiv
	tokMul

	// ( )
	tokLeftParen
	tokRightParen
)

// Helper Methods
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isOp(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '/' || ch == '*'
}

func isNum(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func getOpTok(ch rune) Token {
	switch {
	case ch == '+':
		return tokAdd
	case ch == '-':
		return tokSub
	case ch == '/':
		return tokDiv
	case ch == '*':
		return tokMul
	default:
		return tokError
	}
}

type lexer struct {
	input     string
	lastChar  rune
	pos       int
	lastPos   int
	lastWidth int
}

func newLexer(input string) *lexer {
	return &lexer{
		input:    input,
		lastChar: ' ',
	}
}

func (l *lexer) getTok() (Token, string) {
	defer l.rewind()
	// Store lastPos
	l.lastPos = l.pos
	// Reset lastWidth
	l.lastWidth = 0

	// Skip whitespace
	for isSpace(l.lastChar) {
		l.lastChar = l.next()
	}

	// Capture Number
	if isNum(l.lastChar) {
		numStr := string(l.lastChar)
		l.lastChar = l.next()

		for isNum(l.lastChar) {
			numStr += string(l.lastChar)
			l.lastChar = l.next()
		}

		return tokNumber, numStr
	}

	// Capture Op
	if isOp(l.lastChar) {
		op := l.lastChar
		l.lastChar = l.next()
		return getOpTok(op), ""
	}

	return tokEOF, ""

}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		return eof
	}

	r, width := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += width
	l.lastWidth += width

	return r
}

func (l *lexer) advance() {
	l.pos += l.lastWidth
}

func (l *lexer) peek() Token {
	// store last pos
	lastPos := l.pos
	lastWidth := l.lastWidth

	// Skip Current Token
	l.getTok()
	l.advance()

	// Get next token
	tok, _ := l.getTok()

	l.pos = lastPos
	l.lastWidth = lastWidth

	return tok
}

func (l *lexer) rewind() {
	// move pos back to lastPos
	l.pos = l.lastPos
}
