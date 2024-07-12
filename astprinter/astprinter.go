package astprinter

import (
	"fmt"
	"glox/ast"
	"strings"
)

func visitBinary(expr ast.Binary) {
    return parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func visitGroupingExpr(expr ast.Grouping) {
    return parenthesize("group", expr.Expression)
}

func visitLiteralExpr(expr ast.Literal) string {
    if expr.Value == nil {
        return "nil"
    }

    strVal := fmt.Sprintf("%v", expr.Value)
    return strVal
}

func visitUnaryExpr(expr ast.Unary) {
    return parenthesize(expr.Operator.Lexeme, expr.Right)
}

func parenthesize(name string, exprs ...ast.Visitor{}) {
    var builder strings.Builder

    builder.WriteByte('(')
    builder.WriteString(name)
    for _, expr := range exprs {
        builder.WriteByte(' ')
        builder.WriteString(expr.accept())
    }

    return builder.String()
}

