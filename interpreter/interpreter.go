package interpreter

import (
	"fmt"
	"glox/ast"
	"glox/token"
	"strconv"
)

type interpreter struct {}

func (i *interpreter) NewInterpreter() *interpreter {
    return &interpreter{}
}

func (i *interpreter) VisirForLiteral(expr *ast.Literal) interface{} {
    return expr.Value
}

func (i *interpreter) VisitForUnary(expr *ast.Unary) interface{} {
    right := i.evaluate(expr.Right)

    switch expr.Operator.TokenType {
    case token.Minus:
        n, err := i.convertToFloat64(right)
        if err != nil {
            panic("HANDLE THIS LATER")
        }
        return -n 
    case token.Bang:
        return !i.isTruthy(right)
    default:
        return nil
    }
}

func (i *interpreter) isTruthy(val interface{}) bool {
    switch v := val.(type) {
    case nil:
        return false
    case bool:
        return v
    default:
        return true
    }
}

func (i *interpreter) convertToFloat64(val interface{}) (float64, error) {
    switch v := val.(type) {
    case int:
        return float64(v), nil
    case float64:
        return v, nil
    case float32:
        return float64(v), nil 
    case string:
        n, err := strconv.ParseFloat(v, 64)
        if err != nil {
            return -1, err
        }
        return n, nil
    default:
        return -1, fmt.Errorf("Can't convert to float64")
    }
}

func (i *interpreter) VisirForGrouping(expr *ast.Grouping) interface{} {
    return i.evaluate(expr.Expression)
}

func (i *interpreter) evaluate(expr ast.Expression) interface{} {
    val, _ := expr.Accept(i)
    return val
}
