package msgfmt

import (
	"errors"
)

type literalScanner string

func (scanner literalScanner) Scan(input []byte, kv []interface{}) (int, error) {
	for i := 0; i < len(scanner); i++ {
		if scanner[i] != input[i] {
			return 0, errors.New("failed to match literal: " + string(scanner))
		}
	}
	return 0, nil
}
