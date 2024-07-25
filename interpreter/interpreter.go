package interpreter

import (
	"fmt"
	"glox/ast"
	"glox/token"
	"strconv"
	"strings"
)

type interpreter struct {
	HadRuntimeError bool
}

type runtimeError struct {
	token token.Token
	msg   string
}

func NewInterpreter() *interpreter {
	return &interpreter{HadRuntimeError: false}
}

func (i *interpreter) Interpret(expr ast.Expression) {
	val := i.evaluate(expr)
	if _, ok := val.(runtimeError); ok == true {
		i.HadRuntimeError = true
		err := val.(runtimeError)
		fmt.Printf("%s\n[line %d]\n", err.msg, err.token.Line)
	}

	fmt.Println(i.stringify(val))
}

func (i *interpreter) stringify(obj interface{}) string {
	switch v := obj.(type) {
	case nil:
		return "null"
	case float64:
		txt := strconv.FormatFloat(v, 'E', -1, 64)
		if txt[len(txt)-2:len(txt)-1] == ".0" {
			txt = txt[:len(txt)-2]
			return txt
		}

		return txt
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (i *interpreter) evaluate(expr ast.Expression) interface{} {
	val, _ := expr.Accept(i)
	return val
}

func (i *interpreter) VisitForLiteral(expr *ast.Literal) interface{} {
	return expr.Value
}

func (i *interpreter) VisitForUnary(expr *ast.Unary) interface{} {
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case token.Minus:
		n, err := i.convertToFloat64(right)
		if err != nil {
			return runtimeError{token: expr.Operator, msg: "Operand must be a number"}
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

func (i *interpreter) VisitForGrouping(expr *ast.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *interpreter) VisitForBinary(expr *ast.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	l, lErr := i.convertToFloat64(left)
	r, rErr := i.convertToFloat64(right)

	hasConvErr := lErr != nil && rErr != nil
	if hasConvErr && expr.Operator.TokenType == token.Plus {
		strL := fmt.Sprintf("%v", left)
		strR := fmt.Sprintf("%v", right)

		var builder strings.Builder
		builder.WriteString(strL)
		builder.WriteString(strR)

		return builder.String()
	}

	if hasConvErr && expr.Operator.TokenType != token.Plus {
		// add error handling afterwards
		return runtimeError{token: expr.Operator, msg: "Operand must be a number"}
	}

	switch expr.Operator.TokenType {
	case token.Minus:
		return l - r
	case token.Slash:
		return l / r
	case token.Star:
		return l * r
	case token.Plus:
		return l + r
	case token.Greater:
		return l > r
	case token.GreaterEqual:
		return l >= r
	case token.Less:
		return l < r
	case token.LessEqual:
		return l <= r
	case token.BangEqual:
		return !(l == r)
	case token.EqualEqual:
		return l == r
	}

	return nil
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
