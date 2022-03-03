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

	encoder := jsonutil.NewStdlibStreamEncoder(out, true)
	for _, row := range data {
		err := encoder.EncodeRow(row)
		if err != nil{
			panic(err)
		}
	}

	err := encoder.Close()
	if err != nil {
		panic(err)
	}
}
