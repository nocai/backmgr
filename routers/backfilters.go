package routers

import (
	"backmgr/controllers/backcontrollers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	var SessionFilter = func(ctx *context.Context) {
		requestUrl := ctx.Request.RequestURI
		if requestUrl == "/back/login" || requestUrl == "/back/index.html" {
			return
		}

		adminId := ctx.Input.Session(backcontrollers.ADMIN_ID)
		if adminId == nil {
			beego.Debug("未登录")
			ctx.Redirect(302, "/back/index.html")
		}

	}

	_ = SessionFilter

	// beego.InsertFilter("/back/*", beego.BeforeRouter, SessionFilter)
}
