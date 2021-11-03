package service

import (
	"Gwen/global"
	"Gwen/model"
)

type AdminService struct {
}

func (s *AdminService) Info(id int) *model.Admin {
	a := &model.Admin{}
	global.DB.Where("id = ?", id).First(a)
	return a
}

func (s *AdminService) List() (res *model.AdminListRes) {
	res = &model.AdminListRes{}
	global.DB.Scopes(model.Paginate(1, 10)).Preload("Role").Find(&res.Admins)
	global.DB.Model(&model.Admin{}).Count(&res.TotalSize)
	return
}

func (s AdminService) CheckToken(t string) bool {

	return false
}
