package routers

import (
	"backmgr/controllers/backcontrollers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&backcontrollers.ErrorController{})

	beego.Router("/", &backcontrollers.MainController{})

	beego.Router("/back/index.html", &backcontrollers.IndexController{}, "*:Index")
	beego.Router("/back/main.html", &backcontrollers.IndexController{}, "*:Main")
	beego.Router("/back/left.html", &backcontrollers.IndexController{}, "*:Left")
	beego.Router("/back/welcome.html", &backcontrollers.IndexController{}, "*:Welcome")
	beego.Router("/back/top.html", &backcontrollers.IndexController{}, "*:Top")

	beego.Router("/back/login", &backcontrollers.LoginController{}, "post:Login")

	beego.Router("/back/role/home.html", &backcontrollers.RoleController{}, "*:Home")
	beego.Router("/back/role/add", &backcontrollers.RoleController{}, "post:Add")

}
