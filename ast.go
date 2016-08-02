package main

type AST interface {
	Eval() int
}

type NumAST struct {
	val int
}

func (n *NumAST) Eval() int {
	return n.val
}

type OpAST struct {
	op  Token
	lhs AST
	rhs AST
}

func (n *OpAST) Eval() int {
	switch {
	case n.op == tokSub:
		return n.lhs.Eval() - n.rhs.Eval()
	case n.op == tokAdd:
		return n.lhs.Eval() + n.rhs.Eval()
	case n.op == tokDiv:
		return n.lhs.Eval() / n.rhs.Eval()
	case n.op == tokMul:
		return n.lhs.Eval() * n.rhs.Eval()
	}

	// TODO - Real error handling
	panic("Invalid Op")

}
