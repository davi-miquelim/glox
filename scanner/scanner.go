package scanner

import (
    "glox/token"
    "glox/tokentype"
)

var start := 0
var current := 0
var line := 1

type scanner struct {
    source string
    tokens []interface{}
}

func NewScanner(source string, tokens []string) *scanner {
    s := scanner{source: source, tokens: tokens}
    return &s
}

func scanTokens(s *scanner) []interface{} {
    for isAtEnd != true {
        start = current
        s.scanToken()
    }

    s.tokens = append(s.tokens, token.NewToken(tokentype.Eof, "", nil, line))
    return s.tokens
}

func scanToken(s *scanner) {
    c := s.advance()

    switch c {
    case '(':
        s.addToken(tokentype.LeftParen)
    case ')':
        s.addToken(tokentype.RightParen)
    case '{':
        s.addToken(tokentype.LeftBrace)
    case '}':
        s.addToken(tokentype.RightBrace)
    case ',':
        s.addToken(tokentype.Comma)
    case '.':
        s.addToken(tokentype.Dot)
    case '-':
        s.addToken(tokentype.Minus)
    case '+':
        s.addToken(tokentype.Plus)
    case ';':
        s.addToken(tokenType.SemiColon)
    case '*':
        s.addToken(tokenType.Star)
    }
}

func isAtEnd(s *scanner) boolean {
    return current >= len(s.source)
}

func advance() string {
}
    

