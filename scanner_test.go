package msgfmt

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	"github.com/modern-go/test/must"
)

func TestSscanf(t *testing.T) {
	t.Run("no variable", test.Case(func(ctx context.Context) {
		err := Sscanf("hello", "hello")
		must.Nil(err)
		err = Sscanf("hel", "hello")
		must.NotNil(err)
	}))
}
