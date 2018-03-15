package formatter

import (
	"github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func newJsonFormatter(position int, sampleVal interface{}) Formatter {
	cfg := jsoniter.ConfigDefault
	encoder := cfg.EncoderOf(reflect2.TypeOf(sampleVal))
	encoderKey := reflect2.RTypeOf(sampleVal)
	return FuncFormatter(func(space []byte, kv []interface{}) []byte {
		if position >= len(kv) {
			return space
		}
		stream := cfg.BorrowStream(nil)
		val := kv[position]
		ptr := reflect2.PtrOf(val)
		if reflect2.RTypeOf(val) != encoderKey {
			return formatterOf(position, val).Format(space, kv)
		}
		encoder.Encode(ptr, stream)
		output := append(space, stream.Buffer()...)
		cfg.ReturnStream(stream)
		return output
	})
}
