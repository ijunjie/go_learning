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
	err := r.GetStruct(&data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	err2 := user.SignUp(&data.SignUpInput)
	if err2 != nil {
		response.JsonExit(r, 1, err2.Error())
	}
	response.JsonExit(r, 0, "ok")

}
