package formatter

import (
	"github.com/modern-go/msgfmt/parser"
	"github.com/modern-go/parse"
)

type Formatter interface {
	Format(space []byte, kv []interface{}) []byte
}

func Of(formatStr string, sample []interface{}) Formatter {
	lexer := newFormatterLexer()
	src := parse.NewSourceString(formatStr)
	src.Attachment = sample
	formatter := lexer.Parse(src, 0).(Formatter)
	return formatter
}

func newFormatterLexer() *parser.Lexer {
	return parser.NewLexer(func(l *parser.Lexer) {
		l.ParseLiteral = parseLiteral
		l.ParseVariable = parseVariable
	})
}

func parseLiteral(src *parse.Source, str string) interface{} {
	return literal(str)
}

func parseVariable(src *parse.Source, id string) interface{} {
	return nil
}
