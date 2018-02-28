package test

import (
	"testing"
	"github.com/modern-go/msgfmt/jsonfmt"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"context"
)

func Test_bytes(t *testing.T) {
	t.Run("simple", test.Case(func(ctx context.Context) {
		must.Equal(`"hello"`, jsonfmt.MarshalToString([]byte("hello")))
	}))
	t.Run("unicode", test.Case(func(ctx context.Context) {
		must.Equal(`"\xe4\xb8\xad\xe6\x96\x87"`, jsonfmt.MarshalToString([]byte("中文")))
	}))
	t.Run("unicode and control char", test.Case(func(ctx context.Context) {
		must.Equal(`"\xe4\xb8\xad\n\xe6\x96\x87"`, jsonfmt.MarshalToString([]byte("中\n文")))
	}))
}
