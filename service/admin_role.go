package service

import (
	"Gwen/global"
	"Gwen/model"
	"gorm.io/gorm"
)

type AdminRoleService struct {
}

func (s *AdminRoleService) InfoById(id uint) *model.AdminRole {
	a := &model.AdminRole{}
	global.DB.Where("id = ?", id).First(a)
	return a
}

func (s *AdminRoleService) List(page, pageSize uint, where func(tx *gorm.DB)) (res *model.AdminRoleListRes) {
	res = &model.AdminRoleListRes{}
	tx := global.DB.Model(&model.AdminRole{})
	if where != nil {
		where(tx)
	}
	tx.Count(&res.Total)
	tx.Scopes(Paginate(page, pageSize))
	tx.Find(&res.AdminRoles).Count(&res.Total)
	return
}

func (s *AdminRoleService) Update(ar *model.AdminRole, v map[string]interface{}) error {
	err := global.DB.Model(ar).Updates(v).Error
	return err
}

func (s *AdminRoleService) Create(ar *model.AdminRole) error {
	res := global.DB.Create(ar).Error
	return res
}

func (s *AdminRoleService) Delete(ar *model.AdminRole) error {
	return global.DB.Delete(ar).Error
}
