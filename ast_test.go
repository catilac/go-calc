package main

import "testing"

func TestNumAST(t *testing.T) {
	val := 15
	numAST := &NumAST{val: val}
	if val != numAST.Eval() {
		t.Errorf("numAST.Eval() = %d, expected %d", numAST.Eval(), val)
	}
}

func TestOpAST(t *testing.T) {
	addOpAST := &OpAST{op: tokAdd, lhs: &NumAST{125}, rhs: &NumAST{5}}
	if addOpAST.Eval() != 130 {
		t.Errorf("addOpAST.Eval() == %d, expected %d", addOpAST.Eval(), 130)
	}

	root := &OpAST{op: tokMul, lhs: &NumAST{2}, rhs: addOpAST}
	if root.Eval() != 260 {
		t.Errorf("root.Eval() == %d, expected %d", root.Eval(), 260)
	}

	root = &OpAST{op: tokMul, rhs: &NumAST{2}, lhs: addOpAST}
	if root.Eval() != 260 {
		t.Errorf("root.Eval() == %d, expected %d", root.Eval(), 260)
	}

	subOpAST := &OpAST{op: tokSub, lhs: &NumAST{25}, rhs: &NumAST{5}}
	compound := &OpAST{op: tokMul, lhs: addOpAST, rhs: subOpAST}
	if compound.Eval() != 2600 {
		t.Errorf("compound.Eval() == %d, expected %d", compound.Eval(), 2600)
	}

}
