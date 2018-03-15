package ast

import "github.com/modern-go/parse"

type Literal string

type Variable string

type Func struct {
	Id       string
	FuncName string
	FuncArgs []string
}

type Token interface{}

type Format []interface{}

func ParseLiteral(src *parse.Source, literal string) interface{} {
	return Literal(literal)
}

func ParseVariable(src *parse.Source, id string) interface{} {
	return Variable(id)
}

func ParseFunc(src *parse.Source, id string, funcName string, funcArgs []string) interface{} {
	return Func{id, funcName, funcArgs}
}

func Merge(left interface{}, right interface{}) interface{} {
	leftFormat, _ := left.(Format)
	if leftFormat != nil {
		return append(leftFormat, right)
	}
	return Format{left, right}
}
