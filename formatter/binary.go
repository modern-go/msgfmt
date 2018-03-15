package formatter

import (
	"unicode"
	"unicode/utf8"
)

const hex = "0123456789abcdef"

func newBinaryFormatter(position int) Formatter {
	return FuncFormatter(func(space []byte, kv []interface{}) []byte {
		if int(position) >= len(kv) {
			return space
		}
		val, _ := kv[position].([]byte)
		if val == nil {
			return formatterOf(int(position), kv[position]).Format(space, kv)
		}
		for len(val) > 0 {
			r, n := utf8.DecodeRune(val)
			if r == utf8.RuneError || !unicode.IsPrint(r) {
				for _, b := range val[:n] {
					space = append(space, '\\', 'x', hex[b>>4], hex[b&0xF])
				}
				val = val[n:]
				continue
			}
			space = append(space, val[:n]...)
			val = val[n:]
		}
		return space
	})
}
