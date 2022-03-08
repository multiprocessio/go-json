package jsonutil

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	test := []interface{}{
		map[string]interface{}{"a": float64(1), "b": float64(2)},
		map[string]interface{}{"a": float64(5), "c": float64(3), "d": "xyz"},
	}

	encoders := []func(io.Writer, interface{}) error{EncodeStdlib, Encode}
	for _, encoder := range encoders {
		err := encoder(buf, test)
		assert.Nil(t, err)

		var a interface{}
		err = json.Unmarshal(buf.Bytes(), &a)
		assert.Nil(t, err)
		assert.Equal(t, test, a)

		buf.Reset()
	}
}

func TestEncode_handlesfallback(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	test := []interface{}{
		float64(2),
		float64(1),
	}

	encoders := []func(io.Writer, interface{}) error{EncodeStdlib, Encode}
	for _, encoder := range encoders {
		err := encoder(buf, test)
		assert.Nil(t, err)

		var a interface{}
		err = json.Unmarshal(buf.Bytes(), &a)
		assert.Nil(t, err)
		assert.Equal(t, test, a)

		buf.Reset()
	}
}

func TestEncode_handlesNonObjects(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	test := float64(1)

	encoders := []func(io.Writer, interface{}) error{EncodeStdlib, Encode}
	for _, encoder := range encoders {
		err := encoder(buf, test)
		assert.Nil(t, err)

		var a interface{}
		err = json.Unmarshal(buf.Bytes(), &a)
		assert.Nil(t, err)
		assert.Equal(t, test, a)

		buf.Reset()
	}
}

func TestEncode_handlesNoWrites(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	test := []interface{}{}

	encoders := []func(io.Writer, interface{}) error{EncodeStdlib, Encode}
	for _, encoder := range encoders {
		err := encoder(buf, test)
		assert.Nil(t, err)

		var a interface{}
		err = json.Unmarshal(buf.Bytes(), &a)
		assert.Nil(t, err)
		assert.Equal(t, test, a)

		buf.Reset()
	}
}
