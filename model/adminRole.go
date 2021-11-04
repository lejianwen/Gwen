package model

type AdminRole struct {
	Model
	SeeCb  int      `json:"see_cb"`
	Name   string   `json:"name"`
	Admins []*Admin `gorm:"foreignKey:RoleId" json:"admins,omitempty"`
}
