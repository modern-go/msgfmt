package formatter

import (
	"github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type jsonFormatter struct {
	position int
	encoder jsoniter.ValEncoder
	pool jsoniter.StreamPool
}

func (formatter *jsonFormatter) Format(space []byte, kv []interface{}) []byte {
	stream := formatter.pool.BorrowStream(nil)
	ptr := reflect2.PtrOf(kv[formatter.position])
	formatter.encoder.Encode(ptr, stream)
	output := append(space, stream.Buffer()...)
	formatter.pool.ReturnStream(stream)
	return output
}