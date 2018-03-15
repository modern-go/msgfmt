package formatter

func newStringFormatter(position int) Formatter {
	return FuncFormatter(func(space []byte, kv []interface{}) []byte {
		if int(position) >= len(kv) {
			return space
		}
		val, isString := kv[position].(string)
		if !isString {
			return formatterOf(int(position), kv[position]).Format(space, kv)
		}
		return append(space, val...)
	})
}
