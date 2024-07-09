package token

import (
    "fmt"
)

type Token struct {
    tokenType int
    line int
    lexeme  string
    literal interface{}
}

func NewToken(tokenType int, lexeme string, literal *interface{}, line int) *Token {
    if literal == nil {
        t := Token{lexeme: lexeme, tokenType: tokenType, line: line}
        return &t
    }

    t := Token{lexeme: lexeme, literal: *literal, tokenType: tokenType, line: line}
    return &t
}

func (t *Token) ToString() string {
    str := fmt.Sprintf("%d %s %s", t.tokenType, t.lexeme, t.literal)
    return str
}

