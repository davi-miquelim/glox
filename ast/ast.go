package ast

import "glox/token"

type Visitor interface {
    VisitorForGrouping(Grouping)
    VisitorForLiteral(Literal)
    VisitorForBinary(Binary)
    VisitorForUnary(Unary)
}

type Expression struct {
    *Literal
    *Grouping
    *Binary
}

type Literal struct {
	value interface{}
}

func (obj *Literal) accept(v Visitor) {
    if obj == nil {
        panic("nil Literal")
    }

    v.VisitorForLiteral(*obj)
}

type Grouping struct {
	Expression
}

func (obj *Grouping) accept(v Visitor) {
    if obj == nil {
        panic("nil Grouping")
    }

    v.VisitorForGrouping(*obj)
}

type Binary struct {
	left     Expression
	right    Expression
	operator token.Token
}

func (obj *Binary) accept(v Visitor) {
    if obj == nil {
        panic("nil Grouping")
    }

    v.VisitorForBinary(*obj)
}

type Unary struct {
	right     Expression
	operator token.Token
}

func (obj *Unary) accept(v Visitor) {
    if obj == nil {
        panic("nil Grouping")
    }

    v.VisitorForUnary(*obj)
}
