package stmt

import (
	"errors"
	"glox/ast"
)

type Visitor interface {
	VisitForPrintStmt(*Print)
}

type Stmt struct {
	*Print
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

func (p *Print) Accept(v Visitor) {
	v.VisitForPrintStmt(p)
}
