package main

import (
	"fmt"
	"strconv"
)

type parser struct {
	l    *lexer
	root *AST
}

// input: 2 + 3 * 4 + 5
// output: 19
// if we see a number we know we need to parse a binary op afterwards.
// when we parse the binary op, we know there shoudl be an expression ahead

/*
parse_expression ()
	 return parse_expression_1 (parse_primary (), 0)

parse_expression_1 (lhs, min_precedence)
	lookahead := peek next token

	while lookahead is a binary operator whose precedence is >= min_precedence
		op := lookahead
		advance to next token
		rhs := parse_primary ()
		lookahead := peek next token
		while lookahead is a binary operator whose precedence is greater
			than op's, or a right-associative operator
			whose precedence is equal to op's
			rhs := parse_expression_1 (rhs, lookahead's precedence)
			lookahead := peek next token
		lhs := the result of applying op with operands lhs and rhs

	return lhs
*/

func (p *parser) parseExpression() int {
	return p.parseExpression_1(p.parsePrimary(), 0).Eval()
}

func (p *parser) parsePrimary() AST {
	_, val := p.l.getTok()
	intVal, _ := strconv.Atoi(val)
	return &NumAST{val: intVal}
}

func (p *parser) parseExpression_1(lhs AST, minPrec int) AST {
	lookahead := p.l.peek()
	p.l.advance()

	fmt.Printf("POS: %d, WIDTH: %d, Lookahead: %d\n", p.l.pos, p.l.lastWidth-1, lookahead)

	for getTokPrecedence(lookahead) >= minPrec {
		op := lookahead
		p.l.advance()

		fmt.Printf("POS: %d, WIDTH: %d, Lookahead: %d\n", p.l.pos, p.l.lastWidth-1, lookahead)

		rhs := p.parsePrimary()
		lookahead = p.l.peek()

		fmt.Printf("POS: %d, WIDTH: %d, Lookahead: %d\n", p.l.pos, p.l.lastWidth-1, lookahead)

		for getTokPrecedence(lookahead) > getTokPrecedence(op) {
			rhs = p.parseExpression_1(rhs, getTokPrecedence(lookahead))
			lookahead = p.l.peek()
		}

		node := &OpAST{op: op, lhs: lhs, rhs: rhs}

		lhs = &NumAST{val: node.Eval()}
	}

	return lhs
}

func newParser(input string) *parser {
	return &parser{
		l: newLexer(input),
	}
}

// Helpers

func getTokPrecedence(tok Token) int {
	switch {
	case tok == tokSub:
		return 1
	case tok == tokAdd:
		return 2
	case tok == tokDiv:
		return 3
	case tok == tokMul:
		return 4
	default:
		return -1
	}
}
