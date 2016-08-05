package backmodels

type RoleResourceRef struct {
	Id     int64
	RoleId int64 `orm:"fk"`
	ResId  int64 `orm:"fk"`
}

func (this *RoleResourceRef) TableName() string {
	return "rele_res"
}
