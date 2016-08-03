package back

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "devel:devel@tcp(139.196.191.164:3306)/test?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(Admin))
	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
