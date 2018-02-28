package msgfmt

import (
	"github.com/modern-go/parse"
	"github.com/modern-go/parse/read"
)

type intScanner int

func (scanner intScanner) Scan(src *parse.Source, kv []interface{}) int {
	*kv[scanner].(*int) = int(read.Int64(src))
	return 1
}
