package scanner

import (
    "glox/lox"
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
    case "(":
        s.addToken(tokentype.LeftParen)
    case ")":
        s.addToken(tokentype.RightParen)
    case "{":
        s.addToken(tokentype.LeftBrace)
    case "}":
        s.addToken(tokentype.RightBrace)
    case ",":
        s.addToken(tokentype.Comma)
    case ".":
        s.addToken(tokentype.Dot)
    case "-":
        s.addToken(tokentype.Minus)
    case "+":
        s.addToken(tokentype.Plus)
    case ";":
        s.addToken(tokentype.SemiColon)
    case "*":
        s.addToken(tokentype.Star)
    case "!" && s.match('=') == true:
        s.addToken(tokentype.BangEqual)
    case "!" && s.match('=') == false:
        s.addToken(tokentype.Bang)
    case "=" && s.match('=') == true:
        s.addToken(tokentype.EqualEqual)
    case "=" && s.match('=') == false:
        s.addToken(tokentype.Equal)
    case "<" && s.match('=') == true:
        s.addToken(tokentype.LessEqual)
    case "<" && s.match('=') == false:
        s.addToken(tokentype.Less)
    case ">" && s.match('=') == true:
        s.addToken(tokentype.GreaterEqual)
    case ">" && s.match('=') == false:
        s.addToken(tokentype.Greater)
    case "/" && s.match('/') == true:
        for peek() != "\n" && s.isAtEnd() == false {
            s.advance()
        }
    case "/" && s.match('/') == false:
        s.addToken(tokentype.Slash)
    case " ":
    case "\r":
    case "\t":
    case "\n":
        line++
    case "\"":
        s.str()
    default:
        lox.Error(line, "Unexpected character.")
    }
}

func str(s* scanner) {
    for s.peek() != "\"" && s.isAtEnd() == false {
        if s.peek() == "\n" {
            line++
        }

        s.advance()
    }

    if s.isAtEnd() {
        lox.Error(line, "Unterminated string")
        return 
    }

    s.advance()
    value := s.source[start + 1:current]
    s.addToken(tokentype.String, value)
}

func match(s* scanner, expected string) boolean {
    if s.isAtEnd() {
        return false
    }

    if string(s.source[current]) != expected {
        return false
    }

    current++
    return true
}

func peek(s* scanner) string {
    if s.isAtEnd() {
        return "\0"
    }

    return string(s.source[current])
}

func isAtEnd(s *scanner) boolean {
    return current >= len(s.source)
}

func advance(s *scanner) string {
    return string(s.source[current++])
}

func addToken(s* scanner, type int, literal *interface{}) {
    text := s.source[start:current]

    if literal == nil {
        s.tokens = append(s.tokens, token.NewToken(text, nil, line))
        return
    }

    s.tokens = append(s.tokens, token.NewToken(text, literal, line))
}

