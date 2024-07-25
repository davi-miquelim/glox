package interpreter_test

import (
	"glox/interpreter"
	"glox/parser"
	"glox/token"
	"testing"
)

func TestBinaryOp(t *testing.T) {
    var n interface{} = 1
    var n2 interface{} = 2
    left := token.NewToken(token.Number, "1", &n, 0)
    right := token.NewToken(token.Number, "1", &n2, 0)
    op := token.NewToken(token.Plus, "+", nil, 0)
	eof := token.NewToken(token.Eof, "", nil, 0)
    input := []token.Token{*left, *op, *right, *eof}

    loxParser := parser.NewParser(input)
    exprTree := loxParser.Parse()
    loxInterpreter := interpreter.NewInterpreter()
    loxInterpreter.Interpret(exprTree)
}
