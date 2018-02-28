package test

import (
	"testing"
	"fmt"
	"time"
	. "github.com/modern-go/test"
	. "github.com/modern-go/test/must"
	"github.com/modern-go/msgfmt"
	"context"
)

func Test_time(t *testing.T) {
	epoch := time.Unix(0, 0)
	t.Run("fmt.Sprintf", Case(func(ctx context.Context) {
		Equal("1970-01-01 08:00:00 +0800 CST", fmt.Sprintf("%v", epoch))
	}))
	t.Run("msgfmt.Sprintf", Case(func(ctx context.Context) {
		Equal("Thu Jan  1 08:00:00 1970", msgfmt.Sprintf(
			`{epoch, goTime, Mon Jan _2 15:04:05 2006}`,
			"epoch", epoch))
	}))
	t.Run("msgfmt.Sprintf with default format", Case(func(ctx context.Context) {
		Equal("1970-01-01T08:00:00+08:00", msgfmt.Sprintf(
			`{epoch, goTime}`,
			"epoch", epoch))
	}))
}
