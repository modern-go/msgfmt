package scanner

import (
	"regexp"
	"errors"
)

type Scanner interface {
	Scan(input []byte, kv []interface{}) error
}

func Of(format string, sample []interface{}) Scanner {
	re, err := regexp.Compile(format)
	if err != nil {
		return invalid{err}
	}
	return &rootScanner{
		re: re,
	}
}

type invalid struct {
	err error
}

func (scanner invalid) Scan(input []byte, kv []interface{}) error {
	return scanner.err
}

type rootScanner struct {
	re *regexp.Regexp
}

func (scanner *rootScanner) Scan(input []byte, kv []interface{}) error {
	subMatches := scanner.re.FindSubmatch(input)
	if subMatches == nil {
		return errors.New("input does not match format")
	}
	return nil
}