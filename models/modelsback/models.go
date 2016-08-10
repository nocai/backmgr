package modelsback

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// orm.RegisterDataBase("default", "mysql", "devel:devel@tcp(139.196.191.164:3306)/test?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", "root:root@/backmgr?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(Admin), new(Role), new(Resource), new(RoleResourceRef))
	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
