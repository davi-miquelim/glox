package ast

import (
    "glox/token"
    "errors"
)

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
    *Unary
}

func (obj *Expression) Accept(v Visitor) (error) {
    if obj.Literal != nil {
        v.VisitorForLiteral(*obj.Literal)
    } else if obj.Grouping != nil {
        v.VisitorForGrouping(*obj.Grouping)
    } else if obj.Binary != nil {
        v.VisitorForBinary(*obj.Binary)
    } else if obj.Unary != nil {
        v.VisitorForUnary(*obj.Unary)
    } else {
        return errors.New("nil expression")
    }

    return nil
}


type Literal struct {
	Value interface{}
}

func (obj *Literal) Accept(v Visitor) {
    if obj == nil {
        panic("nil Literal")
    }

    v.VisitorForLiteral(*obj)
}

type Grouping struct {
	Expression
}

func (obj *Grouping) Accept(v Visitor) {
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

func (obj *Binary) Accept(v Visitor) {
    if obj == nil {
        panic("nil Binary")
    }

    v.VisitorForBinary(*obj)
}

type Unary struct {
	Right     Expression
	Operator token.Token
}

func (obj *Unary) Accept(v Visitor) {
    if obj == nil {
        panic("nil Unary")
    }

    v.VisitorForUnary(*obj)
}
