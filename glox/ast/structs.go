package ast

import . "glox/glox/tokens"

type Binary struct {
	Left     interface{}
	Operator Token
	Right    interface{}
}

func (obj *Binary) accept(v AstVisitor) {
	v.VisitBinary(obj)
}

type Grouping struct {
	Expr interface{}
}

func (obj *Grouping) accept(v AstVisitor) {
	v.VisitGrouping(obj)
}

type Literal struct {
	Value interface{}
}

func (obj *Literal) accept(v AstVisitor) {
	v.VisitLiteral(obj)
}

type Unary struct {
	Operator Token
	Right    interface{}
}

func (obj *Unary) accept(v AstVisitor) {
	v.VisitUnary(obj)
}

type AstVisitor interface {
	VisitBinary(x *Binary)
	VisitGrouping(x *Grouping)
	VisitLiteral(x *Literal)
	VisitUnary(x *Unary)
}
