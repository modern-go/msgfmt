package test

import (
	"testing"
	"github.com/modern-go/msgfmt/jsonfmt"
	"errors"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"context"
)

type testError struct {
	onePtr *int
}

func (err testError) Error() string {
	return "hello"
}

func Test_error(t *testing.T) {
	t.Run("ptr struct", test.Case(func(ctx context.Context) {
		must.Equal(`"hello"`, jsonfmt.MarshalToString(errors.New("hello")))
	}))
	t.Run("single ptr struct", test.Case(func(ctx context.Context) {
		type TestObject struct {
			Err testError
		}
		must.Equal(`{"Err":"hello"}`, jsonfmt.MarshalToString(TestObject{}))
	}))
}
