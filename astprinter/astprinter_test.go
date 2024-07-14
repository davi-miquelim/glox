package astprinter_test

import (
	"glox/ast"
	"glox/token"
	"testing"
    "glox/astprinter"
)

func TestAstPrinter(t *testing.T) string  {
    minusToken := token.NewToken(token.Minus, "-", nil, 1)
    minusLiteral := minusToken.Literal
    unary := ast.Unary{Right: ast.Expression{Literal: &ast.Literal{Value: minusLiteral}}}
    star := token.NewToken(token.Star, "*", nil, 1)
    literal := ast.Literal{Value: 45.67}
    grouping := ast.Grouping{Expression: ast.Expression{Literal: &literal}}
    expr := ast.Binary{Left: ast.Expression{Unary: &unary}, Operator: *star, Right: ast.Expression{Grouping: &grouping}}
 
    printer := astprinter.NewPrettyPrinter()

    return "foo"
}
