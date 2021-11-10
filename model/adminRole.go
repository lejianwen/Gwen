package model

type AdminRole struct {
	Model
	Id     uint     `json:"id"`
	SeeCb  uint     `json:"see_cb"`
	Name   string   `json:"name"`
	Admins []*Admin `gorm:"foreignKey:RoleId" json:"admins,omitempty"`
}

type AdminRoleList struct {
	AdminRoles []*AdminRole `json:"list,omitempty"`
}
type AdminRoleListRes struct {
	AdminRoleList
	Pagination
}
