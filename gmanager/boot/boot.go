package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 用于配置初始化.
func init() {
	glog.Info("########service start...")

	v := g.View()
	c := g.Config()
	s := g.Server()

	// 添加配置文件路径
	c.AddPath("config")

	// 添加模板文件路径
	v.AddPath("template")

	// 设置标准输出
	glog.SetStdoutPrint(true)

	// 设置名称转换
	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)

	glog.Info("########service finish.")
}
