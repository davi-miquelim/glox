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
	Value interface{}
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
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (obj *Binary) accept(v Visitor) {
    if obj == nil {
        panic("nil Binary")
    }

    v.VisitorForBinary(*obj)
}

type Unary struct {
	Right     Expression
	Operator token.Token
}

func (obj *Unary) accept(v Visitor) {
    if obj == nil {
        panic("nil Unary")
    }

    v.VisitorForUnary(*obj)
}
