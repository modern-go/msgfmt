package msgfmt_test

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	"github.com/modern-go/test/must"
	"github.com/modern-go/msgfmt"
)

func TestFormatter(t *testing.T) {
	t.Run("literal", test.Case(func(ctx context.Context) {
		must.Equal("hello", msgfmt.Sprintf("hello", nil))
	}))
	t.Run("variable of string", test.Case(func(ctx context.Context) {
		output := msgfmt.Sprintf("hello {var}", "var", "world")
		must.Equal("hello world", output)
	}))
	t.Run("variables", test.Case(func(ctx context.Context) {
		output := msgfmt.Sprintf("hello {var1}{var2}",
			"var1", "world",
			"var2", "!")
		must.Equal("hello world!", output)
	}))
}
