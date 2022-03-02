package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func TestFunction() {

}

func main() {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(TestFunction).Pointer()).Name())
	fmt.Println(reflect.TypeOf(TestFunction).NumIn())
}
