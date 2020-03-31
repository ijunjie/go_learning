package middleware

import (
	"gfdemo/app/service/user"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {
	if user.IsSignedIn(r.Session) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
