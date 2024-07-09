package scanner

import (
	"fmt"
	"glox/token"
	"strconv"
	"unicode"
)

var start = 0
var current = 0
var line = 1
var HadError = false

var keywords = map[string]int{
	"and":    token.And,
	"class":  token.Class,
	"else":   token.Else,
	"false":  token.False,
	"for":    token.For,
	"fun":    token.Fun,
	"if":     token.If,
	"null":   token.Null,
	"or":     token.Or,
	"print":  token.Print,
	"return": token.Return,
	"super":  token.Super,
	"tis":    token.This,
	"true":   token.True,
	"var":    token.Var,
	"while":  token.While,
}

type scanner struct {
	source string
	Tokens []interface{}
}

func NewScanner(source string, tokens []interface{}) *scanner {
	if tokens == nil {
		var emtptyTokens []interface{}
		s := scanner{source: source, Tokens: emtptyTokens}
		return &s
	}

	s := scanner{source: source, Tokens: tokens}
	return &s
}

func (s *scanner) ScanTokens() []interface{} {
	for s.isAtEnd() != true {
		start = current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, token.NewToken(token.Eof, "", nil, line))
	return s.Tokens
}

func (s *scanner) scanToken() {
	c := s.advance()

	switch string(c) {
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
		if unicode.IsDigit(rune(c)) {
			s.number()
		} else if unicode.IsLetter(rune(c)) || string(c) == "_" {
			s.identifier()
		} else {
			Error(line, nil, "Unexpected character.")
		}
	}
}

func Error(line int, where *string, message string) {
	HadError = true
	if where == nil {
		err := fmt.Errorf("[line %d] Error %s", line, message)
		fmt.Println(err)
		return
	}

	err := fmt.Errorf("[line %d] Error %s: %s", line, message, *where)
	fmt.Println(err)
}

func (s *scanner) identifier() {
	c := s.peek()
	for len(c) == 1 && (unicode.IsLetter(rune(c[0])) || unicode.IsDigit(rune(c[0]))) {
		s.advance()
	}

	text := s.source[start : current+1]
	tokenType := keywords[text]

	if tokenType == 0 {
		s.addToken(token.Identifier, nil)
		return
	}

	s.addToken(tokenType, nil)
}

func (s *scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == string(".") && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	strDigit := string(s.source[start:current])
	digit, err := strconv.ParseFloat(strDigit, 64)
	var iDigit interface{} = digit

	if err != nil {
		panic(err)
	}

	s.addToken(token.Number, &iDigit)
}

func (s *scanner) str() {
	for s.peek() != "\"" && s.isAtEnd() == false {
		if s.peek() == "\n" {
			line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		Error(line, nil, "Unterminated string")
		return
	}

	s.advance()
	value := s.source[start+1 : current-1]
	var iValue interface{} = value
	s.addToken(token.String, &iValue)
}

func (s *scanner) match(expected string) bool {
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

func (s *scanner) peekNext() string {
	if current+1 >= len(s.source) {
		return "\\0"
	}

	return string(s.source[current+1])
}

func (s *scanner) isDigit(c string) bool {
	_, err := strconv.Atoi(string(c))

	if err == nil {
		return true
	}

	return false
}

func (s *scanner) isAtEnd() bool {
	return current >= len(s.source)
}

func (s *scanner) advance() byte {
	currentChar := s.source[current]
	current++
	return currentChar
}

func (s *scanner) addToken(tokentype int, literal *interface{}) {
	text := s.source[start:current]

	if literal == nil {
		s.Tokens = append(s.Tokens, token.NewToken(tokentype, text, nil, line))
		return
	}

	s.Tokens = append(s.Tokens, token.NewToken(tokentype, text, literal, line))
}
