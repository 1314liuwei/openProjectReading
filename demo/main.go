package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
)

type Demo struct {
	Name string
	Age  int
}

func main() {
	val := &Demo{Name: "demo", Age: 123}
	rv := reflect.ValueOf(val)
	g.Dump(rv.Kind())

	i := 1
	v := reflect.ValueOf(&i)
	g.Dump(v.Elem().Kind())
}
