package model

import "time"

type Admin struct {
	Model
	RoleId          uint       `json:"role_id"`
	Username        string     `json:"username"`
	Password        string     `json:"-"`
	Nickname        string     `json:"nickname"`
	Token           string     `json:"token,omitempty"`
	TokenExpireTime time.Time  `json:"token_expire_time,omitempty"`
	Status          StatusCode `json:"status"`
	Role            *AdminRole `json:"role,omitempty"`
}
type AdminList struct {
	Admins []*Admin `json:"list,omitempty"`
}
type AdminListRes struct {
	AdminList
	Pagination
}
