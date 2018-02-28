package msgfmt

import (
	"io"
	"github.com/modern-go/parse"
	"fmt"
	"github.com/modern-go/msgfmt/jsonfmt"
	"github.com/modern-go/reflect2"
	"github.com/modern-go/concurrent"
)

type Formatter interface {
	Format(space []byte, kv []interface{}) []byte
}

type Formatters struct {
	formatters []Formatter
}

func (formatters Formatters) Append(formatter Formatter) Formatters {
	return Formatters{
		append(formatters.formatters, formatter),
	}
}

func (formatters Formatters) Format(space []byte, kv []interface{}) []byte {
	for _, formatter := range formatters.formatters {
		space = formatter.Format(space, kv)
	}
	return space
}

var toFormatter = newLexer(func(l *lexer) {
	l.parseVariable = func(src *parse.Source, id string) interface{} {
		sample := src.Attachment.([]interface{})
		idx := findValueIndex(sample, id)
		if idx == -1 {
			src.ReportError(fmt.Errorf("%s not found in args", id))
			return nil
		}
		sampleValue := sample[idx]
		stringer, _ := sampleValue.(fmt.Stringer)
		if stringer != nil {
			return stringerFormatter(idx)
		}
		switch sampleValue.(type) {
		case string:
			return strFormatter(idx)
		case []byte:
			return bytesFormatter(idx)
		default:
			return &jsonFormatter{
				idx:     idx,
				encoder: jsonfmt.EncoderOf(reflect2.TypeOf(sampleValue)),
			}
		}
	}
	l.parseFunc = func(src *parse.Source, id string, funcName string, funcArgs []string) interface{} {
		sample := src.Attachment.([]interface{})
		formatter, err := newFuncFormatter(id, funcName, funcArgs, sample)
		if err != nil {
			src.ReportError(err)
			return nil
		}
		return formatter
	}
	l.parseLiteral = func(src *parse.Source, literal string) interface{} {
		return literalFormatter(literal)
	}
	l.merge = func(left interface{}, right interface{}) interface{} {
		formatters, isFormatters := left.(Formatters)
		if isFormatters {
			return formatters.Append(right.(Formatter))
		}
		return Formatters{[]Formatter{left.(Formatter), right.(Formatter)}}
	}
})

func findValueIndex(sample []interface{}, target string) int {
	for i := 0; i < len(sample); i += 2 {
		key := sample[i].(string)
		if key == target {
			return i + 1
		}
	}
	return -1
}

var formatterCache = concurrent.NewMap()

func FormatterOf(format string, sample []interface{}) Formatter {
	formatterObj, found := formatterCache.Load(format)
	if found {
		return formatterObj.(Formatter)
	}
	formatter := formatterOf(format, sample)
	formatterCache.Store(format, formatter)
	return formatter
}

func formatterOf(format string, sample []interface{}) Formatter {
	src := parse.NewSourceString(format)
	src.Attachment = sample
	formatter := toFormatter.Parse(src, 0)
	if src.Error() != nil {
		if src.Error() == io.EOF {
			return formatter.(Formatter)
		}
		return invalidFormatter(src.Error().Error())
	}
	return invalidFormatter("format not parsed completely")
}
