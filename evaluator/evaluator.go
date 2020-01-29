package evaluator

import (
	"interpreter-in-go/ast"
	"interpreter-in-go/object"
)

var (
	TRUE = &object.Boolean{Value:true}
	FALSE = &object.Boolean{Value:false}
	NULL = &object.Null{}
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input == true {
		return TRUE
	}
	return FALSE
}

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value:node.Value}
	case *ast.Program:
		return EvalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}