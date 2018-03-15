package scanner

import (
	"regexp"
	"errors"
	"github.com/modern-go/msgfmt/parser"
	"github.com/modern-go/parse"
	"fmt"
	"io"
)

type Scanner interface {
	Scan(input []byte, kv []interface{}) error
}

type buildingContext struct {
	fullRe []byte
	sample []interface{}
}

func Of(format string, sample []interface{}) (Scanner, error) {
	src := parse.NewSourceString(format)
	ctx := &buildingContext{
		sample: sample,
	}
	src.Attachment = ctx
	lexer := newScannerLexer()
	scannerObj := lexer.Parse(src, 0)
	if src.Error() != nil && src.Error() != io.EOF {
		return nil, src.Error()
	}
	scanner, _ := scannerObj.(*scanners)
	if scannerObj == nil {
		scanner = &scanners{subScanners: []Scanner{}}
	} else if scanner == nil {
		scanner = &scanners{subScanners: []Scanner{
			scannerObj.(Scanner),
		}}
	}
	re, err := regexp.Compile(string(ctx.fullRe))
	if err != nil {
		return nil, err
	}
	scanner.re = re
	return scanner, nil
}

type FuncScanner func(input []byte, kv []interface{}) error

func (f FuncScanner) Scan(input []byte, kv []interface{}) error {
	return f(input, kv)
}

func newScannerLexer() *parser.Lexer {
	return parser.NewLexer(func(l *parser.Lexer) {
		l.ParseLiteral = parseLiteral
		l.ParseVariable = parseVariable
		l.Merge = merge
	})
}

func merge(left interface{}, right interface{}) interface{} {
	if right == nil {
		return left
	}
	leftScanners, _ := left.(*scanners)
	if leftScanners == nil {
		return &scanners{subScanners: []Scanner{
			left.(Scanner),
			right.(Scanner),
		}}
	}
	leftScanners.subScanners = append(leftScanners.subScanners, right.(Scanner))
	return leftScanners
}

func parseVariable(src *parse.Source, id string) interface{} {
	ctx := src.Attachment.(*buildingContext)
	position := findKey(ctx.sample, id)
	if position == -1 {
		src.ReportError(fmt.Errorf("%s not found", id))
		return nil
	}
	val := ctx.sample[position]
	switch val.(type) {
	case *string:
		return newStringScanner(ctx, position)
	default:
		panic("not implemented")
	}
}

type scanners struct {
	re          *regexp.Regexp
	subScanners []Scanner
}

func (scanner *scanners) Scan(input []byte, kv []interface{}) error {
	indices := scanner.re.FindSubmatchIndex(input)
	if indices == nil {
		return errors.New("input does not match format")
	}
	for i, subScanner := range scanner.subScanners {
		leftIndex := indices[2*(i+1)]
		rightIndex := indices[2*(i+1)+1]
		err := subScanner.Scan(input[leftIndex:rightIndex], kv)
		if err != nil {
			return err
		}
	}
	return nil
}

func findKey(kv []interface{}, target string) int {
	for i := 0; i < len(kv); i += 2 {
		if kv[i].(string) == target {
			return i + 1
		}
	}
	return -1
}
