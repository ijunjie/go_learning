package user

import (
	"gfdemo/app/service/user"
	"gfdemo/library/response"

	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

type SignUpRequest struct {
	user.SignUpInput
}

func (c *Controller) SignUp(r *ghttp.Request) {
	var data SignUpRequest

	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err2 := user.SignUp(&data.SignUpInput); err2 != nil {
		response.JsonExit(r, 1, err2.Error())
	}
	response.JsonExit(r, 0, "ok")

}

type SignInRequest struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

func (c *Controller) SignIn(r *ghttp.Request) {
	var data SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

func (c *Controller) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.IsSignedIn(r.Session))
}