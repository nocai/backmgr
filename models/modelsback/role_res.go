package modelsback

type RoleResourceRef struct {
	Id     int64
	RoleId int64 `orm:"fk"`
	ResId  int64 `orm:"fk"`
}

// 自定义表名 rele_res_ref
func (this *RoleResourceRef) TableName() string {
	return "role_res_ref"
}
