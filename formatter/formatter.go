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
		l.Merge = merge
	})
}

func merge(left interface{}, right interface{}) interface{} {
	leftAsFormatters, _ := left.(formatters)
	if leftAsFormatters != nil {
		return append(leftAsFormatters, right.(Formatter))
	}
	return formatters{left.(Formatter), right.(Formatter)}
}

type formatters []Formatter

func (formatters formatters) Format(space []byte, kv []interface{}) []byte {
	for _, formatter := range formatters {
		space = formatter.Format(space, kv)
	}
	return space
}

func parseLiteral(src *parse.Source, str string) interface{} {
	return literal(str)
}

func parseVariable(src *parse.Source, id string) interface{} {
	sample := src.Attachment.([]interface{})
	position := findKey(sample, id)
	if position == -1 {
		panic("not implemented")
	}
	val := sample[position]
	switch val.(type) {
	case string:
		return stringFormatter(position)
	default:
		panic("not implemented")
	}
}

func findKey(kv []interface{}, target string) int {
	for i := 0; i < len(kv); i += 2 {
		if kv[i].(string) == target {
			return i + 1
		}
	}
	return -1
}
