package astprinter_test

import (
	"fmt"
	"glox/ast"
	"glox/astprinter"
	"glox/token"
	"testing"
)

func TestAstPrinter(t *testing.T) {
	minusToken := token.NewToken(token.Minus, "-", nil, 1)
	var iInt interface{} = 123
	intToken := token.NewToken(token.Number, "123", &iInt, 1)
	intLiteral := ast.NewLiteral(intToken.Literal)
	unary := ast.NewUnary(*minusToken, ast.Expression{Literal: intLiteral})
	star := token.NewToken(token.Star, "*", nil, 1)
	floatLiteral := ast.NewLiteral(45.67)
	grouping := ast.NewGrouping(ast.Expression{Literal: floatLiteral})
	binary := ast.NewBinary(ast.Expression{Unary: unary}, *star, ast.Expression{Grouping: grouping})

	printer := astprinter.NewPrettyPrinter()
	res := fmt.Sprintf("%s", printer.VisitForBinary(binary))

	if res != "(* (- 123) (group 45.67))" {
		t.Errorf("AstPrinter = %s; want (* (-123) (group 45.67))", res)
	}
}
