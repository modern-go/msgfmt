package parser_test

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	"github.com/modern-go/parse"
	"github.com/modern-go/msgfmt/parser"
	"github.com/modern-go/test/must"
)

func TestLexer(t *testing.T) {
	t.Run("literal", test.Case(func(ctx context.Context) {
		src := parse.NewSourceString("hello")
		result := parse.Parse(src, parser.NewLexer(func(l *parser.Lexer) {
			l.ParseLiteral = func(src *parse.Source, literal string) interface{} {
				return literal
			}
		}), 0)
		must.Equal("hello", result)
	}))
}