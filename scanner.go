package msgfmt

import (
	"github.com/modern-go/concurrent"
	"io"
	"github.com/modern-go/parse"
	"errors"
	"fmt"
)

type Scanner interface {
	Scan(src *parse.Source, kv []interface{}) int
}

type Scanners struct {
	scanners []Scanner
}

func (scanners Scanners) Append(scanner Scanner) Scanners {
	return Scanners{
		append(scanners.scanners, scanner),
	}
}

func (scanners Scanners) Scan(src *parse.Source, kv []interface{}) int {
	var count int
	for _, scanner := range scanners.scanners {
		count += scanner.Scan(src, kv)
	}
	return count
}

var toScanner = newLexer(func(l *lexer) {
	l.parseLiteral = func(src *parse.Source, literal string) interface{} {
		return literalScanner(literal)
	}
	l.parseVariable = func(src *parse.Source, id string) interface{} {
		sample := src.Attachment.([]interface{})
		idx := findValueIndex(sample, id)
		if idx == -1 {
			src.ReportError(fmt.Errorf("%s not found in args", id))
			return nil
		}
		sampleValue := sample[idx]
		switch sampleValue.(type) {
		case *int:
			return intScanner(idx)
		default:
			panic("not implemented")
		}
	}
	l.merge = func(left interface{}, right interface{}) interface{} {
		scanners, isScanners := left.(Scanners)
		if isScanners {
			return scanners.Append(right.(Scanner))
		}
		return Scanners{[]Scanner{left.(Scanner), right.(Scanner)}}
	}
})

var scannerCache = concurrent.NewMap()

func ScannerOf(format string, sample []interface{}) (Scanner, error) {
	scannerObj, found := scannerCache.Load(format)
	if found {
		return scannerObj.(Scanner), nil
	}
	scanner, err := scannerOf(format, sample)
	if err != nil {
		return nil, err
	}
	scannerCache.Store(format, scanner)
	return scanner, nil
}

func scannerOf(format string, sample []interface{}) (Scanner, error) {
	src := parse.NewSourceString(format)
	src.Attachment = sample
	scanner := toScanner.Parse(src, 0)
	if src.Error() != nil {
		if src.Error() == io.EOF {
			return scanner.(Scanner), nil
		}
		return nil, src.Error()
	}
	return nil, errors.New("format not parsed completely")
}
