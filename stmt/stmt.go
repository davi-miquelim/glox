package stmt

import (
	"errors"
	"glox/ast"
	"go/token"
)

type Visitor interface {
	VisitForPrintStmt(*Print)
	VisitForVarStmt(*Var)
}

type Stmt struct {
	*Print
    *Var
    Expression *ast.Expression
}

func (statement *Stmt) Accept(v Visitor) (interface{}, error) {
	if statement.Print != nil {
		v.VisitForPrintStmt(statement.Print) 
        return nil, nil
	} else {
		return nil, errors.New("nil expression")
	}
}

type Print struct {
	Expr ast.Expression
}

type Var struct {
    Name token.Token
    Initializer ast.Expression 
}

func (p *Print) Accept(v Visitor) {
	v.VisitForPrintStmt(p)
}
