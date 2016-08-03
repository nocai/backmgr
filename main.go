package main

import (
	_ "backmgr/filter"
	_ "backmgr/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	// beego.SessionOn = true

	beego.Run()
}
