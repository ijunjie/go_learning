package router

import (
	"gfdemo/app/api/user"
	"gfdemo/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.Middleware(middleware.CORS)
		group.ALL("/user", ctlUser)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)
			group.ALL("/user/profile", ctlUser, "Profile")
			group.ALL("/user/signout", ctlUser, "SignOut")
		})
	})
}
