package parser

import (
	"fmt"
	"glox/ast"
	"glox/lox"
	"glox/token"
)

type parser struct {
    tokens []token.Token
    current int
} 

func (p *parser) expresion()  ast.Binary {
    return p.equality()
}

func (p *parser) equality() ast.Binary {
    expr := p.comparison()
    for p.match(token.BangEqual, token.EqualEqual) {
        operator := p.previous()
        right := p.comparison()
        expr = ast.NewBinary(expr, operator, right)
    }

    return expr
}

func (p *parser) comparison() ast.Binary {
    expr := p.term()

    for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
        operator := p.previous()
        right := p.term()
        expr = ast.NewBinary(expr, operator, right)
    } 

    return expr
}

func (p *parser) term() ast.Binary {
    expr := p.factor()

    for p.match(token.Minus, token.Plus) {
        operator := p.previous()
        right := p.factor()
        expr = ast.NewBinary(expr, operator, right)
    } 

    return expr
}

func (p *parser) factor() ast.Binary {
    expr := p.unary()

    for p.match(token.Slash, token.Star) {
        operator := p.previous()
        right := p.unary()
        expr = ast.NewBinary(expr, operator, right)
    } 

    return expr
}

func (p *parser) unary() ast.Unary {
    for p.match(token.Bang, token.Minus) {
        operator := p.previous()
        right := p.unary()
        return *ast.NewUnary(operator, right)
    } 

    return p.primary()
}

func (p *parser) primary() interface{} {
    if p.match(token.False) {
        return *ast.NewLiteral(false)
    }
    if p.match(token.True) {
        return *ast.NewLiteral(true)
    }
    if p.match(token.Null) {
        return *ast.NewLiteral(nil)
    }
    if p.match(token.Number, token.String) {
        return *ast.NewLiteral(p.previous().Literal)
    }
    if p.match(token.LeftParen) {
        expr := p.expresion()
        p.consume(token.RightParen, "Expect ')' after expression. ")
        return *ast.NewGrouping(expr)
    }

    return nil
}

func (p *parser) match(tknTypes ...int) bool {
    for _, tt := range tknTypes {
        if p.check(tt) == true {
            p.advance()
            return true
        }
    }

    return false
}

func (p *parser) consume(tknType int, message string) (*token.Token, error) {
    if p.check(tknType) {
        advance := p.advance()
        return &advance, nil
    }

    currentToken := p.peek()
    return nil, fmt.Errorf("%v %s",currentToken, message)
}

func (p *parser) error(tkn token.Token, message string) {
    if tkn.TokenType == token.Eof {
        where := " at end"
        lox.Error(tkn.Line, &where, message)
    } else {
        where := fmt.Sprintf("at '%s' %s", tkn.Lexeme, message)
        lox.Error(tkn.Line, &where, message)
    }

    return p.NewParseError()
}

func (p *parser) check(tknType int) bool {
    if p.isAtEnd() == true {
        return false
    }

    return p.peek().TokenType == tknType
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

