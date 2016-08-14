package routers

import (
	"backmgr/controllers/controllersback"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllersback.ErrorController{})

	beego.Router("/", &controllersback.MainController{})

	beego.Router("/back/index.html", &controllersback.IndexController{}, "*:Index")
	beego.Router("/back/main.html", &controllersback.IndexController{}, "*:Main")
	beego.Router("/back/left.html", &controllersback.IndexController{}, "*:Left")
	beego.Router("/back/welcome.html", &controllersback.IndexController{}, "*:Welcome")
	beego.Router("/back/top.html", &controllersback.IndexController{}, "*:Top")

	beego.Router("/back/login", &controllersback.LoginController{}, "post:Login")

	// 角色管理
	beego.Router("/back/roles.html", &controllersback.RoleController{}, "get:ListRole")
	beego.Router("/back/roles", &controllersback.RoleController{}, "post:CreateRole")
	beego.Router("/back/roles/:id:int", &controllersback.RoleController{}, "get:GetRoleById")
	beego.Router("/back/roles/:id:int", &controllersback.RoleController{}, "put:UpdateRoleById")
	beego.Router("/back/roles/:id:int", &controllersback.RoleController{}, "delete:DeleteRoleById")

	beego.Router("/back/roles/add.html", &controllersback.RoleController{}, "get:Add")

	// 资源管理
	beego.Router("/back/res.html", &controllersback.ResourceController{}, "get:ListResource")

}
