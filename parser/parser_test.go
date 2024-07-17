package parser_test

import (
	"fmt"
	"glox/parser"
	"glox/token"
	"testing"
)


func TestParser(t *testing.T) {
    tokens := []token.Token{*token.NewToken(token.EqualEqual, "==", nil, 0)}
    loxParser := parser.NewParser(tokens)
    res := loxParser.Parse()
    fmt.Println(res)

}
