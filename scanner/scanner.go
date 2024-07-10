package scanner

import (
	"fmt"
	"glox/token"
	"strconv"
	"unicode"
)


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
	Tokens *[]token.Token
    HadError bool
    current int
    start int
    line int
}

func NewScanner(source string, tokens *[]token.Token) *scanner {
	if tokens == nil {
        s := scanner{source: source, Tokens: nil, current: 0, start: 0, line: 1, HadError: false}
		return &s
	}

	s := scanner{source: source, Tokens: tokens, current: 0, start: 0, line: 1, HadError: false}
	return &s
}

func (s *scanner) ScanTokens() []token.Token {
	for s.isAtEnd() != true {
		s.start = s.current
		s.scanToken()
	}

    eofToken := token.NewToken(token.Eof, "", nil, s.line)

    if s.Tokens == nil {
        tSlice := []token.Token{*eofToken}
        return tSlice 

    }

    tSlice := append(*s.Tokens, *eofToken) 
	s.Tokens = &tSlice 
	return *s.Tokens
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

		if s.match("*") == true {
            s.blockComment()
			break
		}

		s.addToken(token.Slash, nil)
	case " ":
	case "\r":
	case "\t":
	case "\n":
		s.line++
	case "\"":
		s.str()
	default:
		if unicode.IsDigit(rune(c)) {
			s.number()
		} else if unicode.IsLetter(rune(c)) || string(c) == "_" {
			s.identifier()
		} else {
			s.Error(s.line, nil, "Unexpected character.")
		}
	}
}

func (s *scanner) Error(line int, where *string, message string) {
	s.HadError = true
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
	for s.isAtEnd() == true && len(c) == 1 && (unicode.IsLetter(rune(c[0])) || unicode.IsDigit(rune(c[0]))) {
		s.advance()
	}

	text := s.source[s.start : s.current]
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

	strDigit := string(s.source[s.start:s.current])
	digit, err := strconv.ParseFloat(strDigit, 64)
	var iDigit interface{} = digit

	if err != nil {
		panic(err)
	}

	s.addToken(token.Number, &iDigit)
}

func (s *scanner) str() {
	for s.isAtEnd() == false && s.peek() != "\"" {
		if s.peek() == "\n" {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		s.Error(s.line, nil, "Unterminated string")
		return
	}

	s.advance()
	value := s.source[s.start+1 : s.current-1]
	var iValue interface{} = value
	s.addToken(token.String, &iValue)
}

func (s *scanner) blockComment() {
    openCount := 1

    for s.isAtEnd() && openCount > 0 {
        if s.peek() == "\n" {
            s.line++
        }

        isOpenComment := (s.peek() + s.peekNext()) == "/*"
        if isOpenComment {
            openCount++
        }


        isCloseComment := (s.peek() + s.peekNext()) == "*/"
        if isCloseComment  {
            openCount--
        }

        s.advance()
    }

    if s.isAtEnd() && openCount > 0 {
        s.Error(s.line, nil, "Unterminated comment")
        return
    }

    s.advance()
}

func (s *scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}

	if string(s.source[s.current]) != expected {
		return false
	}

	s.current++
	return true
}

func (s *scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}

	return string(s.source[s.current])
}

func (s *scanner) peekNext() string {
	if s.current + 1 >= len(s.source) {
		return "\\0"
	}

	return string(s.source[s.current + 1])
}

func (s *scanner) isDigit(c string) bool {
	_, err := strconv.Atoi(string(c))

	if err == nil {
		return true
	}

	return false
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) advance() byte {
	currentChar := s.source[s.current]
	s.current++
	return currentChar
}

func (s *scanner) addToken(tokentype int, literal *interface{}) {
	text := s.source[s.start:s.current]

    if s.Tokens == nil && literal == nil {
        tSlice := []token.Token{*token.NewToken(tokentype, text, nil, s.line)}
        s.Tokens = &tSlice
    }

    if s.Tokens == nil && literal != nil {
        tSlice := []token.Token{*token.NewToken(tokentype, text, literal, s.line)}
        s.Tokens = &tSlice
    }

	if literal == nil {
        tSlice := append(*s.Tokens, *token.NewToken(tokentype, text, nil, s.line))
		s.Tokens = &tSlice
		return
	}

    tSlice := append(*s.Tokens, *token.NewToken(tokentype, text, literal, s.line))
	s.Tokens = &tSlice
}
