package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

const envRestart = "RESTART"
const envListenFD = "LISTENFD"

func main() {
	path := gfile.SelfPath()
	g.Dump(path)
	g.Dump(os.Environ())
}
