package tests

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	"github.com/modern-go/msgfmt"
	"github.com/modern-go/test/must"
)

func Test_Sscanf(t *testing.T) {
	t.Run("pure literal", test.Case(func(ctx context.Context) {
		must.Equal(0, must.Call(msgfmt.Sscanf, "hello", "hello")[0])
	}))
	t.Run("int", test.Case(func(ctx context.Context) {
		var i int
		n := must.Call(msgfmt.Sscanf,
			"hello: 5", "hello: {i}", "i", &i)[0]
		must.Equal(1, n)
		must.Equal(5, i)
	}))
}
