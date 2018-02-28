package test

import (
	"testing"
	"github.com/modern-go/msgfmt/jsonfmt"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"encoding/json"
	"time"
	"errors"
	"io"
	"context"
)

func Test_pointer(t *testing.T) {
	t.Run("ptr int", test.Case(func(ctx context.Context) {
		one := 1
		must.Equal("1", jsonfmt.MarshalToString(&one))
	}))
	t.Run("nil", test.Case(func(ctx context.Context) {
		var ptr *int
		must.Equal("null", jsonfmt.MarshalToString(ptr))
	}))
	t.Run("ptr eface", test.Case(func(ctx context.Context) {
		one := interface{}(1)
		must.Equal("1", jsonfmt.MarshalToString(&one))
	}))
	t.Run("ptr marshaler", test.Case(func(ctx context.Context) {
		marshaler := json.Marshaler(time.Time{})
		must.Equal(`"0001-01-01T00:00:00Z"`, jsonfmt.MarshalToString(&marshaler))
	}))
	t.Run("ptr error", test.Case(func(ctx context.Context) {
		err := errors.New("hello")
		must.Equal(`"hello"`, jsonfmt.MarshalToString(&err))
	}))
	t.Run("ptr iface", test.Case(func(ctx context.Context) {
		closer := io.Closer(TestCloser(100))
		must.Equal(`100`, jsonfmt.MarshalToString(&closer))
	}))
}
