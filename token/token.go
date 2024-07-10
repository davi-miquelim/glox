package token

import (
    "fmt"
)

type Token struct {
    TokenType int
    Line int
    Lexeme  string
    Literal interface{}
}

func NewToken(tokenType int, lexeme string, literal *interface{}, line int) *Token {
    if literal == nil {
        t := Token{Lexeme: lexeme, TokenType: tokenType, Line: line}
        return &t
    }

    t := Token{Lexeme: lexeme, Literal: *literal, TokenType: tokenType, Line: line}
    return &t
}

func (t *Token) ToString() string {
    str := fmt.Sprintf("%d %s %s", t.TokenType, t.Lexeme, t.Literal)
    return str
}

