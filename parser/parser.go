package parser

import "glox/token"

type parser struct {
    tokens token.Token
    current int
} 

func NewParser(tokens token.Token) *parser {
    return &parser{tokens: tokens, current: 0}
}
