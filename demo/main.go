package main

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

type Demo struct {
	Name string
	Age  int
}

func main() {
	res := (*ghttp.ClientResponse)(nil)
	var res1 *ghttp.ClientResponse

	fmt.Println(res == res1)
}
