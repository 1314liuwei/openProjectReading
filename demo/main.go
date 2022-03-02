package main

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"strings"
)

func TestFunction() {

}

func parsePattern(pattern string) (domain, method, path string, err error) {
	path = strings.TrimSpace(pattern)
	domain = "default"
	method = "ALL"
	if array, err := gregex.MatchString(`([a-zA-Z]+):(.+)`, pattern); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[2])
		if v := strings.TrimSpace(array[1]); v != "" {
			method = v
		}
	}
	if array, err := gregex.MatchString(`(.+)@([\w\.\-]+)`, path); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[1])
		if v := strings.TrimSpace(array[2]); v != "" {
			domain = v
		}
	}
	if path == "" {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "invalid pattern: URI should not be empty")
	}
	if path != "/" {
		path = strings.TrimRight(path, "/")
	}
	return
}

func main() {
	domain, method, uri, err := parsePattern("/*")
	g.Dump(domain, method, uri, err)
	g.Dump(strings.TrimLeft(uri, "/"))
}
