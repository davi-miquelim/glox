package token

import (
    "fmt"
)

type token struct {
    tokenType int
    line int
    lexeme  string
    literal string
}

func NewToken(tokenType int, lexeme string, literal *string, line int) *token {
    if literal == nil {
        t := token{lexeme: lexeme, tokenType: tokenType, line: line}
        return &t
    }

    t := token{lexeme: lexeme, literal: *literal, tokenType: tokenType, line: line}
    return &t
}

func (t *token) ToString() string {
    str := fmt.Sprintf("%d %s %s", t.tokenType, t.lexeme, t.literal)
    return str
}

