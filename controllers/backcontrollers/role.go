package backcontrollers

import (
	"backmgr/controllers"
	"backmgr/models/backmodels"
	// "github.com/astaxie/beego"
	"backmgr/util"
	"github.com/astaxie/beego/validation"
)

type RoleController struct {
	controllers.BaseController
}

func (this *RoleController) Home() {

	util.NewPager(pageNo, pageSize, total, pageData)
	this.TplName = "back/role/home.html"
}

func (this *RoleController) Add() {
	roleName := this.GetString("roleName")

	// 输入验证
	valid := validation.Validation{}
	if v := valid.Required(roleName, "角色名").Message("不能为空！"); !v.Ok {
		this.PrintErrorMsgValid(v)
		return
	}

	role, err := backmodels.GetRoleByName(roleName)
	if err == nil {
		this.PrintErrorMsg("角色名重复")
		return
	}

	role.RoleName = roleName
	_, err = backmodels.AddRole(role)
	if err != nil {
		this.PrintError()
		return
	}
	this.PrintOk()
}
