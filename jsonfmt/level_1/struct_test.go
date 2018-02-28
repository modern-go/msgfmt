package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/modern-go/msgfmt/jsonfmt"
	"github.com/modern-go/reflect2"
)

func Test_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 int `json:"field_2"`
	}
	encoder := jsonfmt.EncoderOf(reflect2.TypeOf(TestObject{}))
	output := encoder.Encode(nil,nil, reflect2.PtrOf(TestObject{"hello", 100}))
	should.Equal(`{"Field1":"hello","field_2":100}`, string(output))
}
