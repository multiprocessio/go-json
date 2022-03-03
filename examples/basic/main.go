package main

import (
	"os"

	"github.com/multiprocessio/go-json"
)

func main() {
	// Uses stdlib's encoding/json
	data := []interface{}{
		map[string]interface{}{"a": 1, "b": 2},
		map[string]interface{}{"a": 5, "c": 3, "d": "xyz"},
	}

	out := os.Stdout // Can be any io.Writer

	err := jsonutil.EncodeStdlib(out, data)
	if err != nil {
		panic(err)
	}
}
