package json

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	goccy_json "github.com/goccy/go-json"
)

func EncodeGeneric(out io.Writer, obj interface{}, marshalFn func(o interface{}) ([]byte, error)) error {
	a, ok := obj.([]interface{})
	// Fall back to normal encoder
	if !ok {
		bs, err := marshalFn(obj)
		if err != nil {
			return err
		}

		for len(bs) > 0 {
			n, err := out.Write(bs)
			if err != nil {
				return err
			}

			bs = bs[n:]
		}
		return nil
	}

	bo := bytes.NewBuffer(nil)
	_, err := bo.Write([]byte("["))
	if err != nil {
		return err
	}

	quotedColumns := map[string][]byte{}

	for i, row := range a {
		// Write a comma before the current object
		if i > 0 {
			_, err = bo.Write([]byte(",\n"))
			if err != nil {
				return err
			}
		}

		r, ok := row.(map[string]interface{})
		if !ok {
			bs, err := marshalFn(row)
			if err != nil {
				return err
			}

			_, err = bo.Write(bs)
			if err != nil {
				return err
			}
			continue
		}

		_, err := bo.Write([]byte("{"))
		if err != nil {
			return err
		}

		j := -1
		for col, val := range r {
			j += 1

			// Write a comma before the current key-value
			if j > 0 {
				_, err = bo.Write([]byte(","))
				if err != nil {
					return err
				}
			}

			quoted := quotedColumns[col]
			if quoted == nil {
				quoted = []byte(strconv.QuoteToASCII(col) + ":")
				quotedColumns[col] = quoted
			}
			_, err = bo.Write(quoted)
			if err != nil {
				return err
			}

			bs, err := marshalFn(val)
			if err != nil {
				return err
			}

			_, err = bo.Write(bs)
			if err != nil {
				return err
			}
		}

		_, err = bo.Write([]byte("}"))
		if err != nil {
			return err
		}
	}

	_, err = bo.Write([]byte("]"))

	for bo.Len() > 0 {
		_, err := bo.WriteTo(out)
		if err != nil {
			return err
		}
	}

	return err
}

func EncodeStdlib(out io.Writer, obj interface{}) error {
	return EncodeGeneric(out, obj, json.Marshal)
}

func EncodeGoccy(out io.Writer, obj interface{}) error {
	return EncodeGeneric(out, obj, goccy_json.Marshal)
}
