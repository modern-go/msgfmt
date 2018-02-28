package msgfmt

import (
	"errors"
	"github.com/modern-go/parse"
)

type literalScanner string

func (scanner literalScanner) Scan(src *parse.Source, kv []interface{}) int {
	input, err := src.PeekN(len(scanner))
	if err != nil {
		src.ReportError(err)
		return 0
	}
	src.ConsumeN(len(scanner))
	for i := 0; i < len(scanner); i++ {
		if scanner[i] != input[i] {
			src.ReportError(errors.New("failed to match literal: " + string(scanner)))
			return 0
		}
	}
	return 0
}
