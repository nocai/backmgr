package controllers

type RightLayoutController struct {
	BaseController
}

// 设置右侧模版
func (this *RightLayoutController) SetRightLayout(nav []string) {
	this.Data["Nav"] = nav
	this.Layout = "back/right.tpl"
}
