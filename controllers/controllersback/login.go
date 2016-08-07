package controllersback

import (
	"backmgr/controllers"
	"backmgr/models/modelsback"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	controllers.BaseController
}

const (
	// session key "adminId"
	ADMIN_ID = "adminId"
	// session key "adminName"
	ADMIN_NAME = "adminName"
	// session key "auths"
	AUTHS = "auths"
)

// 登录
func (this *LoginController) Login() {
	admin := this.GetString("admin")
	pwd := this.GetString("pwd")
	checkInput(admin, pwd)

	adminId := this.GetSession(ADMIN_ID)
	if adminId != nil {
		this.PrintOkMsg("已经登录")
		return
	}

	// 登录
	adminUser, err := modelsback.Login(admin, pwd)
	if err == modelsback.UsernameNotExist {
		this.PrintErrorMsg("用户名不存在")
		return
	}
	if err == modelsback.PasswordNotMatch {
		this.PrintErrorMsg("密码不正确")
		return
	}

	// 取角色
	// role, err := modelsback.GetRoleByAdminId(adminUser.Id)
	// if err != nil {
	// 	beego.Error(err)
	// 	this.PrintErrorMsg("没有角色")
	// 	return
	// }

	// ress, err := modelsback.GetResourceByRoleId(role.Id)
	// if err != nil {
	// 	beego.Error(err)
	// }

	// setSesstionRess(this, ress)
	setSesseionAdmin(this, adminUser)

	this.PrintOk()
}

// 将用户Resource放到Session里
func setSesstionRess(controller *LoginController, ress []modelsback.Resource) {
	var auths []string
	for k, v := range ress {
		beego.Info(k)
		auths[k] = v.ActionPath
	}
	controller.SetSession(AUTHS, auths)
}

// 将用户信息放到Session里
func setSesseionAdmin(controller *LoginController, admin modelsback.Admin) {
	controller.SetSession(ADMIN_ID, admin.Id)
	controller.SetSession(ADMIN_NAME, admin.Username)
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
