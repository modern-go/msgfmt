package formatter

type invalid string

func (invalid invalid) Format(space []byte, kv []interface{}) []byte {
	return append(append(space, "%INVALID% "...), invalid...)
}