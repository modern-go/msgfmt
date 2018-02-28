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
}