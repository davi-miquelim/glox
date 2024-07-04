package scanner

import (
    "glox/lox"
    "glox/token/token"
    "glox/token/tokentype"
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
        scanToken()
    }

    s.tokens = append(s.tokens, token.NewToken(tokentype.Eof, "", nil, line))
    return s.tokens
}

func scanToken(s *scanner) {
    c := s.advance()

    switch c {
    case "(":
        s.addToken(tokentype.LeftParen, nil)
    case ")":
        s.addToken(tokentype.RightParen, nil)
    case "{":
        s.addToken(tokentype.LeftBrace, nil)
    case "}":
        s.addToken(tokentype.RightBrace, nil)
    case ",":
        s.addToken(tokentype.Comma, nil)
    case ".":
        s.addToken(tokentype.Dot, nil)
    case "-":
        s.addToken(tokentype.Minus, nil)
    case "+":
        s.addToken(tokentype.Plus, nil)
    case ";":
        s.addToken(tokentype.SemiColon, nil)
    case "*":
        s.addToken(tokentype.Star, nil)
    case "!":
        if s.match("=") == false {
            s.addToken(tokentype.Bang, nil)
            break
        }

        s.addToken(tokentype.BangEqual, nil)
    case "=":
        if s.match("=") == false {
            s.addToken(tokentype.Equal, nil)
            break
        }

        s.addToken(tokentype.EqualEqual, nil)
    case "<":
        if s.match("=") == false {
            s.addToken(tokentype.LessEqual, nil)
            break
        }

        s.addToken(tokentype.Less, nil)
    case ">":
        if s.match("=") == false {
            s.addToken(tokentype.Greater, nil)
            break
        }

        s.addToken(tokentype.GreaterEqual, nil)
    case "/":
        if s.match("/") == true {
            for s.peek() != "\n" && s.isAtEnd() == false {
                s.advance()
            }
            break
        } 

        s.addToken(tokentype.Slash, nil)
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
    s.addToken(tokentype.String, &value)
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

func (s *scanner) addToken(tokentype int, literal *string) {
    text := s.source[start:current]

    if literal == nil {
        s.tokens = append(s.tokens, token.NewToken(text, nil, line))
        return
    }

    s.tokens = append(s.tokens, token.NewToken(text, literal, line))
}
