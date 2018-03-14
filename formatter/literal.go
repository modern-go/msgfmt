package formatter

type literal string

func (l literal) Format(space []byte, kv []interface{}) []byte {
	return append(space, l...)
}