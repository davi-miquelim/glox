package scanner

import (
    "glox/lox"
    "glox/token"
)

var start = 0
var current = 0
var line = 1

type scanner struct {
    source string
    tokens []interface{}
}

func NewScanner(source string, tokens []interface{}) *scanner {
    s := scanner{source: source, tokens: tokens}
    return &s
}

func scanTokens(s *scanner) []interface{} {
    for s.isAtEnd() != true {
        start = current
        s.scanToken()
    }

    s.tokens = append(s.tokens, token.NewToken(token.Eof, "", nil, line))
    return s.tokens
}

func (s *scanner) scanToken() {
    c := s.advance()

    switch c {
    case "(":
        s.addToken(token.LeftParen, nil)
    case ")":
        s.addToken(token.RightParen, nil)
    case "{":
        s.addToken(token.LeftBrace, nil)
    case "}":
        s.addToken(token.RightBrace, nil)
    case ",":
        s.addToken(token.Comma, nil)
    case ".":
        s.addToken(token.Dot, nil)
    case "-":
        s.addToken(token.Minus, nil)
    case "+":
        s.addToken(token.Plus, nil)
    case ";":
        s.addToken(token.SemiColon, nil)
    case "*":
        s.addToken(token.Star, nil)
    case "!":
        if s.match("=") == false {
            s.addToken(token.Bang, nil)
            break
        }

        s.addToken(token.BangEqual, nil)
    case "=":
        if s.match("=") == false {
            s.addToken(token.Equal, nil)
            break
        }

        s.addToken(token.EqualEqual, nil)
    case "<":
        if s.match("=") == false {
            s.addToken(token.LessEqual, nil)
            break
        }

        s.addToken(token.Less, nil)
    case ">":
        if s.match("=") == false {
            s.addToken(token.Greater, nil)
            break
        }

        s.addToken(token.GreaterEqual, nil)
    case "/":
        if s.match("/") == true {
            for s.peek() != "\n" && s.isAtEnd() == false {
                s.advance()
            }
            break
        } 

        s.addToken(token.Slash, nil)
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

func (s *scanner) str() {
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
    s.addToken(token.String, &value)
}

func (s *scanner) match(expected string)  bool {
    if s.isAtEnd() {
        return false
    }

    if string(s.source[current]) != expected {
        return false
    }

    current++
    return true
}

func (s *scanner) peek() string {
    if s.isAtEnd() {
        return "\\0"
    }

    return string(s.source[current])
}

func (s *scanner) isAtEnd() bool {
    return current >= len(s.source)
}

func (s *scanner) advance() string {
    current++
    return string(s.source[current])
}

func (s *scanner) addToken(tokentype int, literal *[]string) {
    text := s.source[start:current]

    if literal == nil {
        s.tokens = append(s.tokens, token.NewToken(tokentype, text, nil, line))
        return
    }

    s.tokens = append(s.tokens, token.NewToken(tokentype, text, literal, line))
}
