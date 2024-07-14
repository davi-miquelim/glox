package astprinter

import (
	"fmt"
	"glox/ast"
	"strings"
)

type AstPrinter struct {}

func (ast *AstPrinter) VisitForBinary(expr *ast.Binary) interface{} {
    return ast.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ast *AstPrinter) VisitForGrouping(expr *ast.Grouping) interface{} {
    return ast.parenthesize("group", expr.Expression)
}

func (ast *AstPrinter) VisitForLiteral(expr *ast.Literal) interface{} {
    if expr.Value == nil {
        return "nil"
    }

    strVal := fmt.Sprintf("%v", expr.Value)
    return strVal
}

func (ast *AstPrinter) VisitForUnary(expr *ast.Unary) interface{} {
    return ast.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ast *AstPrinter) parenthesize(name string, exprs ...ast.Expression) interface{} {
    var builder strings.Builder

    builder.WriteByte('(')
    builder.WriteString(name)
    for _, expr := range exprs {
        builder.WriteByte(' ')
        err := expr.Accept(ast)

        if err != nil {
            panic(err)
        }

        // builder.WriteString(expr.Accept(ast))
    }

    return builder.String()
}

