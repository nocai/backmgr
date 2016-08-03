package back

// import (
// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// )

// type Role struct {
// 	Id       int64
// 	RoleName string
// }

// func GetRoleById(id int64) (Role, error) {
// 	o := orm.NewOrm()

// 	var role Role
// 	err := o.Raw("select * from role as r where r.id = ?", id).QueryRow(&role)

// 	return role, err
// }
