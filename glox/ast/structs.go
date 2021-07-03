package ast

import . "glox/glox/tokens"

type Grouping struct {
	Expr interface{}
}

func (obj *Grouping) Accept(v AstVisitor) {
	v.VisitGrouping(obj)
}

type Literal struct {
	Value interface{}
}

func (obj *Literal) Accept(v AstVisitor) {
	v.VisitLiteral(obj)
}

type Unary struct {
	Operator Token
	Right    interface{}
}

func (obj *Unary) Accept(v AstVisitor) {
	v.VisitUnary(obj)
}

type Binary struct {
	Left     interface{}
	Operator Token
	Right    interface{}
}

func (obj *Binary) Accept(v AstVisitor) {
	v.VisitBinary(obj)
}

type AstVisitor interface {
	VisitGrouping(x *Grouping)
	VisitLiteral(x *Literal)
	VisitUnary(x *Unary)
	VisitBinary(x *Binary)
}

type Expr interface {
	Accept(x AstVisitor)
}
