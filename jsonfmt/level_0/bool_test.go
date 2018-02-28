package test

import (
	"testing"
	"github.com/modern-go/msgfmt/jsonfmt"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"context"
)

func Test_bool(t *testing.T) {
	t.Run("true", test.Case(func(ctx context.Context) {
		must.Equal("true", jsonfmt.MarshalToString(true))
	}))
}
