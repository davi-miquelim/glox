package ast

import "glox/token"

type expression struct {
    *literal
    *grouping
    *binary
}

type literal struct {
	value interface{}
}

type grouping struct {
	expression
}

type binary struct {
	left     expression
	right    expression
	operator token.Token
}

type unary struct {
	left     expression
	operator token.Token
}
