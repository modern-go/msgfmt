package formatter

type literalFormatter string

func (l literalFormatter) Format(space []byte, kv []interface{}) []byte {
	return append(space, l...)
}