package msgfmt

type literalFormatter string

func (formatter literalFormatter) Format(space []byte, kv []interface{}) []byte {
	return append(space, formatter...)
}