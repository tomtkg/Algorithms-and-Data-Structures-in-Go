package mylib

import (
	"encoding/json"
	"os"
)

func SetStdin(x interface{}) {
	s, _ := json.Marshal(x)
	r, w, _ := os.Pipe()
	w.Write(s)
	w.Close()
	os.Stdin = r
}
