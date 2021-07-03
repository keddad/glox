package ast

import . "glox/glox/tokens"

type Binary struct {
	Left     interface{}
	Operator Token
	Right    interface{}
}

type Grouping struct {
	Expr interface{}
}

type Literal struct {
	Value interface{}
}

type Unary struct {
	Operator Token
	Right    interface{}
}
