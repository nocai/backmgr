package controllersback

import (
	"backmgr/controllers"
)

type ResourceController struct {
	controllers.RightLayoutController
}

func (this *ResourceController) ListResource() {
	this.SetRightLayout([]string{"资源管理", "资源列表"})
	this.TplName = "back/res/home.html"
}
