package main

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"time"
)

type Demo struct {
	Name string
	Age  int
}

func main() {
	ctx := gctx.New()
	p, _ := garray.NewIntArray(true).PopRand()
	s := g.Server(p)
	s.BindHandler("/", func(r *ghttp.Request) {
		tmpPath := gfile.TempDir(guid.S())
		err := gfile.Mkdir(tmpPath)
		gtest.Assert(err, nil)
		defer gfile.Remove(tmpPath)

		file := r.GetUploadFile("file")
		_, err = file.Save(tmpPath)
		gtest.Assert(err, nil)
		r.Response.Write(
			r.Get("json"),
			gfile.GetContents(gfile.Join(tmpPath, gfile.Basename(file.Filename))),
		)
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	path := gdebug.TestDataPath("upload", "file1.txt")
	data := g.Map{
		"file": "@file:" + path,
		"json": `{"uuid": "luijquiopm", "isRelative": false, "fileName": "test111.xls"}`,
	}
	c := g.Client()
	c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
	c.PostContent(ctx, "/", data)
}
