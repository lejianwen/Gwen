package request

type AdminRoleForm struct {
	Id   uint   `json:"id"`
	Name string `json:"name" validate:"required"`
}

type AdminRoleListQuery struct {
	Page     uint   `form:"page"`
	PageSize uint   `form:"page_size"`
	Name     string `form:"name"`
}
