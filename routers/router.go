package routers

import (
	"backmgr/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.MainController{})

	beego.Router("/back/index.html", &controllers.IndexController{}, "*:Index")
	beego.Router("/back/main.html", &controllers.IndexController{}, "*:Main")
	beego.Router("/back/left.html", &controllers.IndexController{}, "*:Left")
	beego.Router("/back/welcome.html", &controllers.IndexController{}, "*:Welcome")
	beego.Router("/back/top.html", &controllers.IndexController{}, "*:Top")

	beego.Router("/back/login", &controllers.LoginController{}, "post:Login")
}
