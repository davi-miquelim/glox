package parser

import (
	"fmt"
	"glox/ast"
	"glox/stmt"
	"glox/token"
)

type parserError struct{}

type parser struct {
	tokens   []token.Token
	current  int
	HadError bool
}

func (p *parser) Parse() []stmt.Stmt {
    var statements []stmt.Stmt

    for p.isAtEnd() == false {
        statements = append(statements, p.statement())
    }

	return statements
}

func (p *parser) expresion() ast.Expression {
	return p.equality()
}

func (p *parser) statement() stmt.Stmt {
    if p.match(token.Print) {
        return p.printStatement()
    }

    return p.expressionStatement()
}

func (p *parser) printStatement() stmt.Stmt {
    val := p.expresion()
    p.consume(token.SemiColon, "Expect ';' after value.")

    return stmt.Stmt{Print: &stmt.Print{Expr: val} }
}

func (p *parser) expressionStatement() stmt.Stmt {
    expr := p.expresion()
    p.consume(token.SemiColon, "Expect ';' after expression.")

    return stmt.Stmt{Expression: &expr}
}

func (p *parser) equality() ast.Expression {
	expr := p.comparison()
	for p.match(token.BangEqual, token.EqualEqual) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.Expression{Binary: ast.NewBinary(expr, operator, right)}
	}

	return expr
}

func (p *parser) comparison() ast.Expression {
	expr := p.term()

	for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
		operator := p.previous()
		right := p.term()
		expr = ast.Expression{Binary: ast.NewBinary(expr, operator, right)}
	}

	return expr
}

func (p *parser) term() ast.Expression {
	expr := p.factor()

	for p.match(token.Minus, token.Plus) {
		operator := p.previous()
		right := p.factor()
		expr = ast.Expression{Binary: ast.NewBinary(expr, operator, right)}
	}

	return expr
}

func (p *parser) factor() ast.Expression {
	expr := p.unary()

	for p.match(token.Slash, token.Star) {
		operator := p.previous()
		right := p.unary()
		expr = ast.Expression{Binary: ast.NewBinary(expr, operator, right)}
	}

	return expr
}

func (p *parser) unary() ast.Expression {
	for p.match(token.Bang, token.Minus) {
		operator := p.previous()
		right, err := p.primary()
		if err != nil {
            return ast.Expression{}
		}

		unary := ast.NewUnary(operator, ast.Expression{Literal: right.Literal})
		return ast.Expression{Unary: unary}
	}

	primary, err := p.primary()
	if err != nil {
        return ast.Expression{}
	}

	return primary
}

func (p *parser) primary() (ast.Expression, error) {
	if p.match(token.False) {
		literal := *ast.NewLiteral(false)
		return ast.Expression{Literal: &literal}, nil
	}
	if p.match(token.True) {
		literal := *ast.NewLiteral(true)
		return ast.Expression{Literal: &literal}, nil
	}
	if p.match(token.Null) {
		literal := *ast.NewLiteral(nil)
		return ast.Expression{Literal: &literal}, nil
	}
	if p.match(token.Number, token.String) {
		literal := *ast.NewLiteral(p.previous().Literal)
		return ast.Expression{Literal: &literal}, nil
	}
	if p.match(token.LeftParen) {
		expr := p.expresion().Binary
		p.consume(token.RightParen, "Expect ')' after expression. ")
		grouping := *ast.NewGrouping(ast.Expression{Binary: expr})
		return ast.Expression{Grouping: &grouping}, nil
	}

	return ast.Expression{}, fmt.Errorf("%v", p.parseError(p.peek(), "Expected expression"))
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

	currentTkn := p.peek()
	return nil, fmt.Errorf("%v", p.parseError(currentTkn, message))
}

func (p *parser) parseError(tkn token.Token, message string) *parserError {
	if tkn.TokenType == token.Eof {
		where := " at end"
		p.report(tkn.Line, &where, message)
	} else {
		where := fmt.Sprintf("at '%s' %s", tkn.Lexeme, message)
		p.report(tkn.Line, &where, message)
	}

	return newParserError()
}

func (p *parser) report(line int, where *string, message string) error {
	p.HadError = true

	if where == nil {
		err := fmt.Errorf("[line %d] Error %s", line, message)
        fmt.Println(err)
		return err
	}

	err := fmt.Errorf("[line %d] Error %s: %s", line, message, *where)
    fmt.Println(err)
	return err
}

func (p *parser) synchronize() {
	p.advance()
	for p.isAtEnd() == false {
		if p.previous().TokenType == token.SemiColon {
			return
		}

		switch p.peek().TokenType {
		case token.Class:
		case token.For:
		case token.Fun:
		case token.If:
		case token.Print:
		case token.Return:
		case token.Var:
		case token.While:
		}

		p.advance()
	}
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
	return p.tokens[p.current-1]
}

func NewParser(tokens []token.Token) *parser {
	return &parser{tokens: tokens, current: 0, HadError: false}
}

func newParserError() *parserError {
	return &parserError{}
}
