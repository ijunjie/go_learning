package router

import (
	"gfdemo/app/api/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.ALL("/user", ctlUser)
	})
}
