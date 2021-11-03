package model

type Admin struct {
	Model
	RoleId   int        `json:"role_id"`
	Username string     `json:"username"`
	Nickname string     `json:"nickname"`
	Role     *AdminRole `json:"role,omitempty"`
}
type AdminList struct {
	Admins []*Admin `json:"list,omitempty"`
}
type AdminListRes struct {
	AdminList
	Pagination
}
