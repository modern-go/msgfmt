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
}