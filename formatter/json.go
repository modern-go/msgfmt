package formatter

import (
	"github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type jsonFormatter struct {
	position   int
	encoder    jsoniter.ValEncoder
	encoderKey uintptr
	cfg        jsoniter.API
}

func (formatter *jsonFormatter) Format(space []byte, kv []interface{}) []byte {
	stream := formatter.cfg.BorrowStream(nil)
	val := kv[formatter.position]
	ptr := reflect2.PtrOf(val)
	if reflect2.RTypeOf(val) != formatter.encoderKey {
		return formatterOf(formatter.position, val).Format(space, kv)
	}
	formatter.encoder.Encode(ptr, stream)
	output := append(space, stream.Buffer()...)
	formatter.cfg.ReturnStream(stream)
	return output
}
