package request

import (
	"Gwen/model"
)

type AdminForm struct {
	RoleId   uint             `json:"role_id" binding:"required"`
	Username string           `json:"username" binding:"required"`
	Password string           `json:"password" binding:"required"`
	Nickname string           `json:"nickname" binding:"required"`
	Status   model.StatusCode `json:"status" binding:"required"`
}
