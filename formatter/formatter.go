package formatter

import (
	"github.com/modern-go/msgfmt/parser"
	"github.com/modern-go/parse"
	"github.com/modern-go/concurrent"
)

var funcRegistry = concurrent.NewMap()

type Formatter interface {
	Format(space []byte, kv []interface{}) []byte
}

type Func interface {
	FuncName() string
	FormatterOf(funcArgs []string, sample []interface{}, id string) Formatter
}

func RegisterFunc(f Func) {
	funcRegistry.Store(f.FuncName(), f)
}

type FuncFormatter func(space []byte, kv []interface{}) []byte

func (f FuncFormatter) Format(space []byte, kv []interface{}) []byte {
	return f(space, kv)
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
		l.ParseFunc = parseFunc
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
	return literalFormatter(str)
}

func parseFunc(src *parse.Source, id string, funcName string, funcArgs []string) interface{} {
	funcObj, found := funcRegistry.Load(funcName)
	if !found {
		return invalid("func " + funcName + " not registered")
	}
	sample := src.Attachment.([]interface{})
	formatter := funcObj.(Func).FormatterOf(funcArgs, sample, id)
	return formatter
}

func parseVariable(src *parse.Source, id string) interface{} {
	sample := src.Attachment.([]interface{})
	position := findKey(sample, id)
	if position == -1 {
		return invalid(id + " not found in arguments")
	}
	val := sample[position]
	return formatterOf(position, val)
}

func formatterOf(position int, val interface{}) Formatter {
	switch val.(type) {
	case string:
		return newStringFormatter(position)
	case []byte:
		return newBinaryFormatter(position)
	default:
		return newJsonFormatter(position, val)
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
