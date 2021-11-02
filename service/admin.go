package service

import (
	"Gwen/global"
	"Gwen/model"
)

type AdminService struct {
}

func (s *AdminService) Info(id int) model.Admin {
	var a model.Admin
	global.DB.Where("id = ?", id).First(&a)
	return a
}

func (s *AdminService) List() []model.Admin {
	var list []model.Admin
	global.DB.Scopes(model.Paginate(1, 10)).Preload("Role").Find(&list)
	return list
}
