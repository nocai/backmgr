package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

const (
	OK  int = 200
	Bad int = 400

	PageNo   = 1
	PageSize = 10
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
	this.PrintOkMsg("操作成功")
}

func (this *BaseController) PrintOkMsg(msg string) {
	this.PrintOkMsgData(msg, nil)
}

func (this *BaseController) PrintOkMsgData(msg string, data interface{}) {
	this.Print(OK, msg, data)
}

func (this *BaseController) PrintError() {
	this.PrintErrorMsg("系统异常")
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

func (this *BaseController) PrintErrorMsgValid(r *validation.Result) {
	this.PrintErrorMsg(r.Error.Key + " " + r.Error.Message)
}
