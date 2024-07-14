package ast

import (
    "glox/token"
    "errors"
)

type Visitor interface {
    VisitForGrouping(*Grouping) interface{}
    VisitForLiteral(*Literal) interface{}
    VisitForBinary(*Binary) interface{}
    VisitForUnary(*Unary) interface{} 
}

type Expression struct {
    *Literal
    *Grouping
    *Binary
    *Unary
}

func (obj *Expression) Accept(v Visitor) (error) {
    if obj.Literal != nil {
        v.VisitForLiteral(obj.Literal)
    } else if obj.Grouping != nil {
        v.VisitForGrouping(obj.Grouping)
    } else if obj.Binary != nil {
        v.VisitForBinary(obj.Binary)
    } else if obj.Unary != nil {
        v.VisitForUnary(obj.Unary)
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

    v.VisitForLiteral(obj)
}

type Grouping struct {
	Expression
}

func (obj *Grouping) Accept(v Visitor) {
    if obj == nil {
        panic("nil Grouping")
    }

    v.VisitForGrouping(obj)
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

    v.VisitForBinary(obj)
}

type Unary struct {
	Right     Expression
	Operator token.Token
}

func (obj *Unary) Accept(v Visitor) {
    if obj == nil {
        panic("nil Unary")
    }

    v.VisitForUnary(obj)
}
