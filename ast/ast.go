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

func (obj *Expression) Accept(v Visitor) (interface{}, error) {
    if obj.Literal != nil {
        return v.VisitForLiteral(obj.Literal), nil
    } else if obj.Grouping != nil {
        return v.VisitForGrouping(obj.Grouping), nil
    } else if obj.Binary != nil {
        return v.VisitForBinary(obj.Binary), nil
    } else if obj.Unary != nil {
        return v.VisitForUnary(obj.Unary), nil
    } else {
        return nil, errors.New("nil expression")
    }
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
