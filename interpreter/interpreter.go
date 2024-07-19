package interpreter

import "glox/ast"

type interpreter struct {}

func (i *interpreter) NewInterpreter() *interpreter {
    return &interpreter{}
}

func (i *interpreter) VisirForLiteral(expr *ast.Literal) interface{} {
    return expr.Value
}
