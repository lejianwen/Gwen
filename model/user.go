package model

const (
	GAME_NOT_FINISHED uint = 0
	GAME_FINISHED     uint = 1
)

type User struct {
	Model
	Nickname string     `json:"nickname"`
	Avatar   string     `json:"avatar"`
	Openid   string     `json:"openid"`
	Status   StatusCode `json:"status"`
}

type UserList struct {
	Users []*User `json:"list,omitempty"`
}
type UserListRes struct {
	UserList
	Pagination
}
