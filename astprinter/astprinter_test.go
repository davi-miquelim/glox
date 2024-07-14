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
    minusLiteral := minusToken.Literal
    unary := ast.Unary{Right: ast.Expression{Literal: &ast.Literal{Value: minusLiteral}}}
    star := token.NewToken(token.Star, "*", nil, 1)
    literal := ast.Literal{Value: 45.67}
    grouping := ast.Grouping{Expression: ast.Expression{Literal: &literal}}
    binary := ast.Binary{Left: ast.Expression{Unary: &unary}, Operator: *star, Right: ast.Expression{Grouping: &grouping}}
 
    printer := astprinter.NewPrettyPrinter()
    res := fmt.Sprintf("%s", printer.VisitForBinary(&binary))

    if res != "(* (-123) (group 45.67))" {
        t.Errorf("AstPrinter = %s; want (* (-123) (group 45.67))", res)
    }
}