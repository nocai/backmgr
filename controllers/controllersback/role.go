package controllersback

import (
	"backmgr/controllers"
	"backmgr/models/modelsback"
	// "github.com/astaxie/beego"
	// "backmgr/util"
	"github.com/astaxie/beego/validation"
)

type RoleController struct {
	controllers.BaseController
}

func (this *RoleController) Home() {
	// var pageNo int64
	// var pageSize int64
	// var err error

	pageNo, _ := this.GetInt64("pageNo", controllers.PageNo)
	pageSize, _ := this.GetInt64("pageSize", controllers.PageSize)
	roleName := this.GetString("roleName")
	pager := modelsback.PageRole(pageNo, pageSize, roleName)

	this.Data["Pager"] = pager
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

	role, err := modelsback.GetRoleByName(roleName)
	if err == nil {
		this.PrintErrorMsg("角色名重复")
		return
	}

	role.RoleName = roleName
	_, err = modelsback.AddRole(role)
	if err != nil {
		this.PrintError()
		return
	}
	this.PrintOk()
}
