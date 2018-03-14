package parser_test

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	"github.com/modern-go/parse"
	"github.com/modern-go/msgfmt/parser"
	"github.com/modern-go/test/must"
	"github.com/modern-go/msgfmt/parser/ast"
)

func TestLexer(t *testing.T) {
	t.Run("literal", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("hello")
		lexer := parser.NewAstLexer()
		result := lexer.Parse(src, 0)
		must.Equal(ast.Literal("hello"), result)
	}))
	t.Run("variable", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("hello {world}")
		lexer := parser.NewAstLexer()
		result := lexer.Parse(src, 0)
		must.Equal(ast.Format{
			ast.Literal("hello "),
			ast.Variable("world"),
		}, result)
	}))
	t.Run("func", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("hello {world, goTime, Mon Jan _2 15:04:05 2006}")
		lexer := parser.NewAstLexer()
		result := lexer.Parse(src, 0)
		must.Equal(ast.Format{
			ast.Literal("hello "),
			ast.Func{
				"world",
				"goTime",
				[]string{" Mon Jan _2 15:04:05 2006"},
			},
		}, result)
	}))
}