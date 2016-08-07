package controllersback

import (
	"backmgr/controllers"
	"backmgr/models/modelsback"
	"github.com/astaxie/beego"
	// "backmgr/util"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

type RoleController struct {
	controllers.BaseController
}

func (this *RoleController) ListRole() {
	pageNo, _ := this.GetInt("pageNo", controllers.PageNo)
	this.Data["PageNo"] = pageNo

	pageSize, _ := this.GetInt("pageSize", controllers.PageSize)
	this.Data["PageSize"] = pageSize

	roleName := this.GetString("roleName")
	beego.Debug("roleName = " + roleName)
	this.Data["RoleName"] = roleName

	orderBy := this.GetString("orderBy")
	beego.Debug("orderBy = " + orderBy)

	order := this.GetString("order")
	beego.Debug("order = " + order)

	// 默认按添加时间反序
	if len(orderBy) == 0 || len(order) == 0 {
		orderBy = "create_time"
		order = "desc"
	}
	this.Data["OrderBy"] = orderBy
	this.Data["Order"] = order

	pager, err := modelsback.PageRole(pageNo, pageSize, roleName, orderBy, order)
	if err != nil {
		beego.Error(err)
	}

	this.Data["Pager"] = pager
	this.TplName = "back/role/home.html"
}

func (this *RoleController) CreateRole() {
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
	role.CreateTime = time.Now()
	_, err = modelsback.AddRole(role)
	if err != nil {
		beego.Error(err)
		this.PrintError()
		return
	}
	this.PrintOk()
}

func (this *RoleController) GetRoleById() {
	id := this.GetString(":id")
	if len(id) == 0 {
		beego.Error("roleId is nil")
		return
	}

	roleId, _ := strconv.ParseInt(id, 10, 64)
	role, err := modelsback.GetRoleById(roleId)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Role"] = role
	this.TplName = "back/role/role.html"

}

func (this *RoleController) UpdateRoleById() {
	id := this.GetString(":id")
	if len(id) == 0 {
		beego.Error("id is nil")
		return
	}
	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
		return
	}

	roleName := this.GetString("roleName")
	if len(roleName) == 0 {
		beego.Error("roleName is nil")
		this.PrintErrorMsg("roleName 为空")
		return
	}

	role, err := modelsback.GetRoleByName(roleName)
	if err != nil && err != modelsback.RoleNotExist {
		beego.Error(err)
		return
	}

	if role.Id != roleId && role.Id != 0 {
		this.PrintErrorMsg("角色名重复")
		return
	}

	err = modelsback.UpdateRoleNameById(roleName, roleId)
	if err != nil {
		beego.Error(err)
		this.PrintErrorMsg("更新失败")
		return
	}

	this.PrintOk()
}

func (this *RoleController) DeleteRoleById() {
	id := this.GetString(":id")
	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
		this.PrintErrorMsg("id is nil")
		return
	}

	num, err := modelsback.DeleteRoleById(roleId)
	if err != nil {
		beego.Error(err)
		this.PrintErrorMsg("删除失败")
		return
	}
	beego.Info("删除%s条数据", strconv.FormatInt(num, 10))

	this.PrintOk()
}

func (this *RoleController) Add() {
	this.Data["Name"] = this.GetSession(ADMIN_NAME)
	this.TplName = "back/role/role.html"
}
