package controllers

import (
	"backmgr/models/back"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

const (
	// Session Key "adminId"
	ADMINID = "adminId"
)

// 登录
func (this *LoginController) Login() {
	admin := this.GetString("admin")
	pwd := this.GetString("pwd")
	checkInput(admin, pwd)

	adminId := this.GetSession(ADMINID)
	if adminId != nil {
		this.PrintOkMsg("已经登录")
		return
	}

	adminUser, err := back.Login(admin, pwd)
	if err == back.UsernameNotExist {
		this.PrintErrorMsg("用户名不存在")
		return
	}
	if err == back.PasswordNotMatch {
		this.PrintErrorMsg("密码不正确")
		return
	}

	this.SetSession(ADMINID, adminUser.Id)
	this.SetSession("adminName", adminUser.Username)

	this.PrintOk()
}

// 输入验证
func checkInput(admin, pwd string) {
	valid := validation.Validation{}
	if v := valid.Required(admin, "admin"); !v.Ok {
		beego.Info(v.Error.Key, v.Error.Message)
	}
	if v := valid.Required(pwd, "pwd"); !v.Ok {
		beego.Info(v.Error.Key, v.Error.Message)
	}
}
