package formatter

type stringFormatter int

func (position stringFormatter) Format(space []byte, kv []interface{}) []byte {
	val, isString := kv[position].(string)
	if !isString {
		return formatterOf(int(position), kv[position]).Format(space, kv)
	}
	return append(space, val...)
}