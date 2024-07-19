package interpreter

import "glox/ast"

type interpreter struct {}

func (i *interpreter) NewInterpreter() *interpreter {
    return &interpreter{}
}

func (i *interpreter) VisirForLiteral(expr *ast.Literal) interface{} {
    return expr.Value
}

func (i *interpreter) VisirForGrouping(expr *ast.Grouping) interface{} {
    return i.evaluate(expr.Expression)
}

func (i *interpreter) evaluate(expr ast.Expression) interface{} {
    val, _ := expr.Accept(i)
    return val
}
