package scanner

import "errors"

func newStringScanner(ctx *buildingContext, position int) Scanner {
	ctx.fullRe = append(ctx.fullRe, "(.*)"...)
	return FuncScanner(func(input []byte, kv []interface{}) error {
		val, _ := kv[position].(*string)
		if val == nil {
			return errors.New("variable is not string pointer")
		}
		*val = string(input)
		return nil
	})
}