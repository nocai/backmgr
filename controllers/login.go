package controllers

import (
	"backmgr/models/back"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

// 登录
func (this *LoginController) Login() {
	beego.Debug("host = " + this.Ctx.Request.URL.String())
	admin := this.GetString("admin")
	pwd := this.GetString("pwd")

	valid := validation.Validation{}
	if v := valid.Required(admin, "admin"); !v.Ok {
		beego.Info(v.Error.Key, v.Error.Message)
	}
	if v := valid.Required(pwd, "pwd"); !v.Ok {
		beego.Info(v.Error.Key, v.Error.Message)
	}

	adminId := this.GetSession("adminId")
	if adminId != nil {
		this.PrintOkMsg("已经登录")
		return
	}

	adminUser, err := back.GetAdminByUsername(admin)
	if err != nil {
		this.PrintErrorMsg("用户名不存在")
		return
	}
	if adminUser.Password != pwd {
		this.PrintErrorMsg("密码不正确")
		return
	}

	this.SetSession("adminId", adminUser.Id)
	this.SetSession("adminName", adminUser.Username)
	beego.Debug(this.GetSession("adminId"))

	this.PrintOk()
}
