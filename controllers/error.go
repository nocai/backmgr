package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "back/error404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "back/501.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "back/dberror.tpl"
}
