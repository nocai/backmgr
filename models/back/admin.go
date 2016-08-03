package back

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id       int64
	Username string
	Password string
}

func GetAdminByUsername(username string) (Admin, error) {
	o := orm.NewOrm()

	var admin Admin
	err := o.Raw("select * from admin as admin where admin.username = ?", username).QueryRow(&admin)
	return admin, err
}

func ExistOfUsername(username string) bool {
	admin, err := GetAdminByUsername(username)
	if err != nil {
		beego.Info(err)
		return false
	}

	return admin.Id > 0
}
