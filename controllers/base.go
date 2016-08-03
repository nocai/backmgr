package controllers

import (
	"github.com/astaxie/beego"
)

const (
	OK  int = 200
	Bad int = 400
)

type BaseController struct {
	beego.Controller
}

type JsonMsg struct {
	Code int
	Msg  string
	Data interface{}
}

func (this *BaseController) PrintOk() {
	this.PrintOkMsg("success")
}

func (this *BaseController) PrintOkMsg(msg string) {
	this.PrintOkMsgData(msg, nil)
}

func (this *BaseController) PrintOkMsgData(msg string, data interface{}) {
	this.Print(OK, msg, data)
}

func (this *BaseController) PrintError() {
	this.PrintErrorMsg("fail")
}

func (this *BaseController) PrintErrorMsg(msg string) {
	this.PrintErrorMsgData(msg, nil)
}

func (this *BaseController) PrintErrorMsgData(msg string, data interface{}) {
	this.Print(Bad, msg, data)
}

func (this *BaseController) Print(code int, msg string, data interface{}) {
	jsonMsg := &JsonMsg{Code: code, Msg: msg, Data: data}
	beego.Debug(jsonMsg)
	this.Data["json"] = jsonMsg

	this.ServeJSON()
}
