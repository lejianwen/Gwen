package request

import (
	"Gwen/model"
)

type AdminForm struct {
	Id       uint             `json:"id"`
	RoleId   uint             `json:"role_id" validate:"required,gt=0" label:"角色id"`
	Username string           `json:"username" validate:"required"`
	Password string           `json:"password" validate:"required"`
	Nickname string           `json:"nickname" validate:"required"`
	Status   model.StatusCode `json:"status" validate:"required,gte=0"`
}

type AdminListQuery struct {
	Page     uint   `form:"page"`
	PageSize uint   `form:"page_size"`
	Nickname string `form:"nickname"`
}
