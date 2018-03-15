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
	t.Run("variable of integer", test.Case(func(ctx context.Context) {
		output := msgfmt.Sprintf("hello {var}", "var", 100)
		must.Equal("hello 100", output)
	}))
	t.Run("variable type might change", test.Case(func(ctx context.Context) {
		output := msgfmt.Sprintf("hello {var}", "var", 100)
		must.Equal("hello 100", output)
		output = msgfmt.Sprintf("hello {var}", "var", "world")
		must.Equal("hello world", output)
	}))
	t.Run("variable not found", test.Case(func(ctx context.Context) {
		msgfmt.Sprintf("hello {var}")
	}))
	t.Run("variable of bytes", test.Case(func(ctx context.Context) {
		output := msgfmt.Sprintf("hello {var}", "var", []byte("world"))
		must.Equal("hello world", output)
		output = msgfmt.Sprintf("hello {var}", "var", []byte{1,2,3})
		must.Equal(`hello \x01\x02\x03`, output)
		output = msgfmt.Sprintf("hello {var}", "var", []byte{0xc3, 0x28})
		must.Equal(`hello \xc3(`, output)
	}))
}
