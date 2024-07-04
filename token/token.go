package token

import (
    "fmt"
    "glox/tokentype"
)

type token struct {
    type int
    line int
    lexeme  string
    literal interface{}
}

func NewToken(lexeme string, literal interface{}, type, line int) *token {
    t := token{lexeme: lexeme, literal: literal, type: type, line: line}
    return &t
}

func ToString(t *token) {
    return t.type + " " t.lexeme + " " + t.literal
}


