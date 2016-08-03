package routers

import (
	"backmgr/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	var SessionFilter = func(ctx *context.Context) {
		requestUrl := ctx.Request.RequestURI
		if requestUrl == "/back/login" || requestUrl == "/back/index.html" {
			return
		}

		adminId := ctx.Input.Session(controllers.ADMINID)
		if adminId == nil {
			beego.Debug("未登录")
			ctx.Redirect(302, "/back/index.html")
		}

	}

	beego.InsertFilter("/back/*", beego.BeforeRouter, SessionFilter)
}
