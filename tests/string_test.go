package tests

import (
	"testing"
	"time"
	"github.com/modern-go/msgfmt"
	"fmt"
	. "github.com/modern-go/test"
	. "github.com/modern-go/test/must"
	"context"
)

func Test_string(t *testing.T) {
	t.Run("string => string", Case(func(ctx context.Context) {
		AssertEqual("ahellob", fmt.Sprintf("a%sb", "hello"))
		AssertEqual("ahellob", msgfmt.Sprintf("a{key}b", "key", "hello"))
	}))
	t.Run("int => string", Case(func(ctx context.Context) {
		AssertEqual("%!s(int=100)", fmt.Sprintf("%s", 100))
		AssertEqual("100", msgfmt.Sprintf("{key}", "key", 100))
	}))
	t.Run("bytes => string", Case(func(ctx context.Context) {
		AssertEqual("hello", fmt.Sprintf("%s", []byte("hello")))
		AssertEqual("hello", msgfmt.Sprintf("{key}", "key", []byte("hello")))
	}))
	t.Run("printf", Case(func(ctx context.Context) {
		fmt.Printf("%s\n", "hello")
		msgfmt.Printf("{key}\n", "key", "hello")
	}))
	t.Run("println", Case(func(ctx context.Context) {
		fmt.Println("hello", "world")
		msgfmt.Println("hello", "world")
	}))
}

func Benchmark_string_to_string(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		msgfmt.Sprintf("{key}", "key", "hello")
		//fmt.Sprintf("%s", "hello")
	}
}

func Benchmark_time_now(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		time.Now().String()
	}
}
