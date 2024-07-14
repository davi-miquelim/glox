package ast

import (
	"errors"
	"glox/token"
)

type Visitor interface {
	VisitForGrouping(*Grouping) interface{}
	VisitForLiteral(*Literal) interface{}
	VisitForBinary(*Binary) interface{}
	VisitForUnary(*Unary) interface{}
}

type derivations interface {
    *Literal | *Grouping | *Binary | *Unary
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

func NewLiteral(value interface{}) *Literal {
    return &Literal{Value: value}
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

func NewGrouping(expr Expression) *Grouping {
    return &Grouping{Expression: expr}
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

func NewBinary(left Expression, operator token.Token, right Expression) *Binary {
    return &Binary{Left: left, Right: right, Operator: operator}
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

func NewUnary(operator token.Token, right Expression) *Unary {
    return &Unary{Right: right, Operator: operator}
}

func (obj *Unary) Accept(v Visitor) {
	if obj == nil {
		panic("nil Unary")
	}

	v.VisitForUnary(obj)
}
