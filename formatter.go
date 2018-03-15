package msgfmt

import (
	"github.com/modern-go/concurrent"
	"github.com/modern-go/msgfmt/formatter"
)

var formatterCache = concurrent.NewMap()

func Sprintf(format string, kv ...interface{}) string {
	fmtObj, found := formatterCache.Load(format)
	if found {
		return string(fmtObj.(formatter.Formatter).Format(nil, kv))
	}
	fmt := formatter.Of(format, kv)
	formatterCache.Store(format, fmt)
	return string(fmt.Format(nil, kv))
}
