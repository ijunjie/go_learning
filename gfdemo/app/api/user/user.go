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
	var data *SignUpRequest
	err := r.GetStruct(&data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
}
