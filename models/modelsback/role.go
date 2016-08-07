package modelsback

import (
	"backmgr/util"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

var (
	RoleNotExist = errors.New("角色不存在")
)

// 角色
type Role struct {
	Id         int64
	RoleName   string
	CreateTime time.Time
}

// 分页
func PageRole(pageNo, pageSize int, roleName, orderBy, order string) (*util.Pager, error) {
	var countSql = "select count(r.id) from role as r where 1 = 1"
	var pageSql = "select * from role as r where 1 = 1"

	if len(roleName) > 0 {
		countSql += " and r.role_name like '%" + roleName + "%'"
		pageSql += " and r.role_name like '%" + roleName + "%'"
	}

	if len(orderBy) != 0 && len(order) != 0 {
		pageSql += " order by " + orderBy + " " + order
	}

	beego.Debug(pageNo)
	beego.Debug(pageSize)

	if pageNo > 0 && pageSize > 0 {
		var startIndex = (pageNo - 1) * pageSize
		beego.Debug(startIndex)

		pageSql += " limit " + strconv.Itoa(startIndex) + ", " + strconv.Itoa(pageSize)
	}

	o := orm.NewOrm()

	var total int64
	err := o.Raw(countSql).QueryRow(&total)

	var roles []Role
	_, err = o.Raw(pageSql).QueryRows(&roles)
	if err != nil {
		return util.NewPager(pageNo, pageSize, 0, nil), err
	}

	var pageData []interface{}
	for _, role := range roles {
		pageData = append(pageData, role)
	}
	return util.NewPager(pageNo, pageSize, total, pageData), nil
}

func GetRoleById(id int64) (Role, error) {
	o := orm.NewOrm()

	var role Role
	err := o.Raw("select * from role as r where r.id = ?", id).QueryRow(&role)

	return role, err
}

func UpdateRoleNameById(roleName string, roleId int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("update role set role_name = ? where id = ?", roleName, roleId).Exec()
	return err
}

// 删除
// 返回删除的记录数
func DeleteRoleById(roleId int64) (int64, error) {
	o := orm.NewOrm()

	result, err := o.Raw("delete from role where id = ?", roleId).Exec()
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return affected, nil
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

	res, err := o.Raw("insert into role (role_name, create_time) values (?, ?)", role.RoleName, role.CreateTime).Exec()
	if err != nil {
		return Role{}, err
	}

	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return Role{}, err
	}

	return Role{id, role.RoleName, role.CreateTime}, nil
}

func GetRoleByName(roleName string) (Role, error) {
	o := orm.NewOrm()

	var roles []Role
	num, err := o.Raw("select * from role as r where r.role_name = ?", roleName).QueryRows(&roles)
	if err == nil {
		if num > 0 {
			return roles[0], nil
		} else {
			return Role{}, RoleNotExist
		}
	}
	return Role{}, err
}
