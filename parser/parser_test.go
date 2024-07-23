package parser_test

import (
	"fmt"
	"glox/parser"
	"glox/token"
	"testing"
)

func TestParser(t *testing.T) {
	var iNum interface{} = 1
	eof := token.NewToken(token.Eof, "", nil, 0)
	tokens := []token.Token{*token.NewToken(token.Number, "1", &iNum, 0), *token.NewToken(token.EqualEqual, "==", nil, 0), *token.NewToken(token.Number, "1", &iNum, 0), *eof}
	loxParser := parser.NewParser(tokens)
	res := loxParser.Parse()
	fmt.Println(res)
}
