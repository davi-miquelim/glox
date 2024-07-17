package astprinter

import (
	"fmt"
	"glox/ast"
	"strings"
)

type astPrinter struct{}

func NewPrettyPrinter() *astPrinter {
	return &astPrinter{}
}

func (ast *astPrinter) VisitForBinary(expr *ast.Binary) interface{} {
	return ast.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ast *astPrinter) VisitForGrouping(expr *ast.Grouping) interface{} {
	return ast.parenthesize("group", expr.Expression)
}

func (ast *astPrinter) VisitForLiteral(expr *ast.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	strVal := fmt.Sprintf("%v", expr.Value)
	return strVal
}

func (ast *astPrinter) VisitForUnary(expr *ast.Unary) interface{} {
	return ast.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ast *astPrinter) parenthesize(name string, exprs ...ast.Expression) interface{} {
	var builder strings.Builder

	builder.WriteByte('(')
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteByte(' ')
		val, err := expr.Accept(ast)

		if err != nil {
			panic(err)
		}

		strVal := fmt.Sprintf("%v", val)
		builder.WriteString(strVal)
	}
	builder.WriteByte(')')

	return builder.String()
}
