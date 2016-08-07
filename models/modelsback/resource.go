package modelsback

import (
	"github.com/astaxie/beego/orm"
)

type Resource struct {
	Id   int64
	Name string
	// 资源路径
	ActionPath string
	// 页面路径
	PagePath string
}

// 以角色ID取对应资源
func GetResourceByRoleId(roleId int64) ([]Resource, error) {
	o := orm.NewOrm()

	var ress []Resource
	_, err := o.Raw(`select * from resource as res 
		left join rele_res as rr on res.id = rr.res_id
		left join role as r on r.id = rr.role_id
		where r.id = ?
		`, roleId).QueryRows(&ress)
	return ress, err
}
