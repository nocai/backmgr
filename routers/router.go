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

	beego.Router("/back/role/home.html", &controllersback.RoleController{}, "*:Home")
	beego.Router("/back/role/add", &controllersback.RoleController{}, "post:Add")

}
