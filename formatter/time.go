package formatter

import (
	"time"
	"strings"
)

func init() {
	RegisterFunc(&goTime{})
}

type goTime struct {
}

func (f *goTime) FuncName() string {
	return "goTime"
}

func (f *goTime) FormatterOf(funcArgs []string, sample []interface{}, id string) Formatter {
	position := findKey(sample, id)
	if position == -1 {
		return invalid(id + " not found")
	}
	if len(funcArgs) == 0 {
		return invalid("goTime format layout not specified")
	}
	layout := strings.TrimSpace(funcArgs[0])
	return FuncFormatter(func(space []byte, kv []interface{}) []byte {
		if position >= len(kv) {
			return invalid(id + " not found").Format(space, kv)
		}
		val, isTime := kv[position].(time.Time)
		if !isTime {
			return invalid(id + " is not time").Format(space, kv)
		}
		return append(space, val.Format(layout)...)
	})
}

