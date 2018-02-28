package msgfmt

import (
	"github.com/modern-go/concurrent"
	"io"
	"github.com/modern-go/parse"
	"errors"
)

type Scanner interface {
	Scan(input []byte, kv []interface{}) (int, error)
}

var toScanner = newLexer(func(l *lexer) {
	l.parseLiteral = func(src *parse.Source, literal string) interface{} {
		return literalScanner(literal)
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

