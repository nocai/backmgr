package backmodels

import (
	"backmgr/util"
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	RoleNotExist = errors.New("角色不存在")
)

// 角色
type Role struct {
	Id       int64
	RoleName string
}

func pageRole(pageNo, pageSize int64, roleName string) util.Pager {
	var countSql = "select count(r.id) from role as r where 1 = 1"
	var pageSql = "select * from role as r where 1 = 1"

	if len(roleName) > 0 {
		countSql += " and r.role_name = '" + roleName + "'"
		pageSql += " and r.role_name = '" + roleName + "'"
	}

	o := orm.NewOrm()

	r, err := o.Raw(countSql).Exec()
	if err != nil {
		return *util.NewPager(pageNo, pageSize, 0, nil)
	}
	var total int64
	total, err = r.LastInsertId()
	if err != nil {
		return *util.NewPager(pageNo, pageSize, 0, nil)
	}

	var roles []Role
	_, err = o.Raw(pageSql).QueryRows(&roles)
	if err != nil {
		return *util.NewPager(pageNo, pageSize, 0, nil)
	}

	var pageData []interface{}
	for _, role := range roles {
		pageData = append(pageData, role.(interface{}))
	}
	return *util.NewPager(pageNo, pageSize, total, &pageData)
}

func GetRoleById(id int64) (Role, error) {
	o := orm.NewOrm()

	var role Role
	err := o.Raw("select * from role as r where r.id = ?", id).QueryRow(&role)

	return role, err
}

func GetRoleByAdminId(adminId int64) (Role, error) {
	o := orm.NewOrm()

	var roles []Role
	num, err := o.Raw("select * from role as r where r.id = (select id from admin as admin where admin.roleId = ?)", adminId).QueryRows(&roles)
	if err == nil && num > 0 {
		return roles[0], err
	}
	return Role{}, RoleNotExist
}

func AddRole(role Role) (Role, error) {
	o := orm.NewOrm()

	res, err := o.Raw("insert into role (role_name) values (?)", role.RoleName).Exec()
	if err != nil {
		return Role{}, err
	}

	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return Role{}, err
	}

	return Role{id, role.RoleName}, nil
}

func GetRoleByName(roleName string) (Role, error) {
	o := orm.NewOrm()

	var roles []Role
	num, err := o.Raw("select * from role as r where r.role_name = ?", roleName).QueryRows(&roles)
	if err == nil && num > 0 {
		return roles[0], err
	}
	return Role{}, RoleNotExist
}
