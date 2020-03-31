package user

import (
	"errors"
	"fmt"
	"gfdemo/app/model/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

type SignUpInput struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

func SignUp(data *SignUpInput) error {

	if err := gvalid.CheckStruct(data, nil); err != nil {
		return errors.New(err.FirstString())
	}
	if data.Nickname == "" {
		data.Nickname = data.Passport
	}

	if !CheckPassport(data.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Passport))
	}

	if !CheckNickName(data.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data.Nickname))
	}

	var entity user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}

	entity.CreateTime = gtime.Now()
	if _, err := user.Save(entity); err != nil {
		return err
	}
	return nil
}

func CheckPassport(passport string) bool {
	count, err := user.FindCount("passport", passport)
	if err != nil {
		return false
	} else {
		return count == 0
	}
}

func CheckNickName(nickname string) bool {
	if i, err := user.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

func SignIn(passport string, password string, session *ghttp.Session) error {
	one, err := user.FindOne("passport=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("账号或密码错误")
	}
	return session.Set(USER_SESSION_MARK, one)
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	glog.Infof("session: %+v", session.Map())
	return session.Contains(USER_SESSION_MARK)
}

func SignOut(session *ghttp.Session) error {
	return session.Remove(USER_SESSION_MARK)
}

func GetProfile(session *ghttp.Session) (u *user.Entity) {
	var r user.Entity
	_ = session.GetStruct(USER_SESSION_MARK, &r)
	return &r
}
