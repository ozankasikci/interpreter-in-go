package evaluator

import (
	"interpreter-in-go/ast"
	"interpreter-in-go/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value:node.Value}
	case *ast.Program:
		return EvalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}

	return nil
}

func EvalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
