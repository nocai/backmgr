package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	PassUrl = "/back/index.html"
)

func init() {
	var SessionFilter = func(ctx *context.Context) {
		requestUrl := ctx.Request.RequestURI
		if requestUrl == "/back/login" || requestUrl == "/back/index.html" {
			return
		}

		adminId := ctx.Input.Session("adminId")
		if adminId == nil {
			beego.Debug("未登录")
			ctx.Redirect(302, PassUrl)
		}

	}

	beego.InsertFilter("/back/*", beego.BeforeRouter, SessionFilter)
}
