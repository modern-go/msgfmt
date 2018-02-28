package test

import (
	"testing"
	"github.com/modern-go/msgfmt/jsonfmt"
	"time"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"context"
)

func Test_time(t *testing.T) {
	t.Run("epoch", test.Case(func(ctx context.Context) {
		must.Equal(`"0001-01-01T00:00:00Z"`, jsonfmt.MarshalToString(time.Time{}))
	}))
}
