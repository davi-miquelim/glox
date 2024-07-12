package ast

import "glox/token"

type visitor interface {
    visitorForGrouping(grouping)
    visitorForLiteral(literal)
    visitorForBinary(binary)
    visitorForUnary(unary)
}

type expression struct {
    *literal
    *grouping
    *binary
}

type literal struct {
	value interface{}
}

func (obj *literal) accept(v visitor) {
    if obj == nil {
        panic("nil literal")
    }

    v.visitorForLiteral(*obj)
}

type grouping struct {
	expression
}

func (obj *grouping) accept(v visitor) {
    if obj == nil {
        panic("nil grouping")
    }

    v.visitorForGrouping(*obj)
}

type binary struct {
	left     expression
	right    expression
	operator token.Token
}

func (obj *binary) accept(v visitor) {
    if obj == nil {
        panic("nil grouping")
    }

    v.visitorForBinary(*obj)
}

type unary struct {
	right     expression
	operator token.Token
}

func (obj *unary) accept(v visitor) {
    if obj == nil {
        panic("nil grouping")
    }

    v.visitorForUnary(*obj)
}
