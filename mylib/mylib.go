package mylib

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func PrintFuncName(f interface{}) {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	arr := strings.Split(name, ".")
	fmt.Println(arr[len(arr)-1])
}

func SetStdin(x interface{}) {
	s, _ := json.Marshal(x)
	r, w, _ := os.Pipe()
	w.Write(s)
	w.Close()
	os.Stdin = r
}
