package controllersback

import (
	"backmgr/controllers"
	// "backmgr/models/modelsback"
	// "github.com/astaxie/beego"
	// "backmgr/util"
	// "github.com/astaxie/beego/validation"
	// "strconv"
	// "time"
)

type ResourceController struct {
	controllers.BaseController
}

func (this *ResourceController) ListResource() {
	this.Layout = "back/layout.html"
	this.TplName = "back/res/home.tpl"
}
