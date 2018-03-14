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
}
