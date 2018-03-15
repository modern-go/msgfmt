package msgfmt

import "github.com/modern-go/msgfmt/scanner"

func Sscanf(str string, format string, kv ...interface{}) error {
	return scanner.Of(format, kv).Scan([]byte(str), kv)
}
