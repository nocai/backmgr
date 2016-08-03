package controllers

import (
// "github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.TplName = "back/login.html"
}

func (this *IndexController) Main() {
	this.TplName = "back/main.html"
}

func (this *IndexController) Top() {
	this.TplName = "back/top.html"
}

func (this *IndexController) Left() {
	this.TplName = "back/left.html"
}

func (this *IndexController) Welcome() {
	this.TplName = "back/welcome.html"
}
