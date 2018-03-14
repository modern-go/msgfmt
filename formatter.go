package msgfmt

import (
	"github.com/modern-go/msgfmt/formatter"
)

func Sprintf(format string, kv ...interface{}) string {
	fmt := formatter.Of(format, kv)
	return string(fmt.Format(nil, kv))
}
