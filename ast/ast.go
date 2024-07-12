package ast

import "glox/token"

type expression struct {
}

type literal struct {
	value interface{}
}

type grouping struct {
	expression
    left token.Token
    right token.Token
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
