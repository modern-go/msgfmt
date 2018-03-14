package formatter

type stringFormatter int

func (position stringFormatter) Format(space []byte, kv []interface{}) []byte {
	val := kv[position].(string)
	return append(space, val...)
}