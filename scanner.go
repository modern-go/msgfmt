package msgfmt

import "github.com/modern-go/msgfmt/scanner"

func Sscanf(str string, format string, kv ...interface{}) error {
	scanner, err := scanner.Of(format, kv)
	if err != nil {
		return err
	}
	return scanner.Scan([]byte(str), kv)
}
