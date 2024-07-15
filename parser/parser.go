package parser

import (
	"glox/token"
	"glox/ast"
)

type parser struct {
    tokens []token.Token
    current int
} 

func (p *parser) expresion() interface{} {
    return p.equality()
}

func (p *parser) equality() {
    expr := p.comparison()
    for p.match(token.BangEqual, token.EqualEqual) {
        operator := p.previous()
        right := p.comparison()
        expr = ast.NewBinary(expr, operator, right)
    }

    return expr
}

func (p *parser) match(tokenTypes ...int) bool {
    for _, tt := range tokenTypes {
        if p.check(tt) == true {
            p.advance()
            return true
        }
    }

    return false
}

func (p *parser) check(tokenType int) bool {
    if p.isAtEnd() == true {
        return false
    }

    return p.peek().TokenType == tokenType
}

func (p *parser) advance() token.Token {
    if p.isAtEnd() == false {
        p.current++
    }

    return p.previous()
}

func (p *parser) isAtEnd() bool {
    currentToken := p.peek()
    if currentToken.TokenType == token.Eof {
        return true 
    }

    return false
}

func (p *parser) peek() token.Token {
    return p.tokens[p.current]
}

func (p *parser) previous() token.Token {
    return p.tokens[p.current - 1]
}

func NewParser(tokens []token.Token) *parser {
    return &parser{tokens: tokens, current: 0}
}

