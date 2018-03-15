package msgfmt

import (
	"context"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"testing"
)

func TestSscanf(t *testing.T) {
	t.Run("no variable", test.Case(func(ctx context.Context) {
		err := Sscanf("hello", "hel")
		must.Nil(err)
		err = Sscanf("hel", "hello")
		must.NotNil(err)
	}))
	t.Run("scan string", test.Case(func(ctx context.Context) {
		var str string
		err := Sscanf("hello world", "hello {var}", "var", &str)
		must.Nil(err)
		must.Equal("world", str)
	}))
	t.Run("merge", test.Case(func(ctx context.Context) {
		var str string
		err := Sscanf("hello world!", "hello {var}!", "var", &str)
		must.Nil(err)
		must.Equal("world", str)
	}))
}
