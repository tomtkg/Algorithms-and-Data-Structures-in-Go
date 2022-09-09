package mylib

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Tuple[T1, T2 any] struct {
	V1 T1
	V2 T2
}

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

func SetString(x string) {
	r, w, _ := os.Pipe()
	w.Write([]byte(x))
	w.Close()
	os.Stdin = r
}
