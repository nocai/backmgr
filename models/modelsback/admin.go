package modelsback

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	UsernameNotExist = errors.New("用户名不存在")
	PasswordNotMatch = errors.New("密码不匹配")
)

type Admin struct {
	Id       int64
	Username string
	Password string
	RoleId   int64
}

// 取角色
func GetRole(admin Admin) (Role, error) {
	return GetRoleByAdminId(admin.Id)
}

// 登录
func Login(username, password string) (Admin, error) {
	admin, err := GetAdminByUsername(username)
	if err != nil {
		return Admin{}, UsernameNotExist
	}

	if admin.Password != password {
		return admin, PasswordNotMatch
	}
	return admin, nil
}

// 取管理员帐号By Usernam
func GetAdminByUsername(username string) (Admin, error) {
	o := orm.NewOrm()

	var admin Admin
	err := o.Raw("select * from admin as admin where admin.username = ?",
		username).QueryRow(&admin)

	return admin, err
}

// Username是否存在
func ExistOfUsername(username string) bool {
	admin, err := GetAdminByUsername(username)
	if err != nil {
		beego.Info(err)
		return false
	}

	return admin.Id > 0
}
