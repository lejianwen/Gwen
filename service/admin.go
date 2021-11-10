package service

import (
	"Gwen/global"
	"Gwen/model"
	"Gwen/model/request"
	"Gwen/utils"
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type AdminService struct {
}

func (s *AdminService) Info(id uint) *model.Admin {
	a := &model.Admin{}
	global.DB.Scopes(CommonEnable()).Where("id = ?", id).First(a)
	return a
}

func (s *AdminService) List(page, pageSize uint, where func(tx *gorm.DB)) (res *model.AdminListRes) {
	res = &model.AdminListRes{}
	tx := global.DB.Model(&model.Admin{})
	tx.Scopes(Paginate(page, pageSize))
	if where != nil {
		where(tx)
	}
	tx.Find(&res.Admins).Count(&res.TotalSize)
	return
}

func (s *AdminService) InfoByPwd(username, password string) *model.Admin {
	a := &model.Admin{}
	enP := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	global.DB.Scopes(CommonEnable()).Where("username = ? and password = ?", username, enP).First(a)
	return a
}

func (s *AdminService) GenerateToken(admin *model.Admin) string {
	rand.Seed(time.Now().UnixNano())
	// id + pwd
	str := strconv.FormatUint(uint64(admin.Id), 10) + admin.Password + string(rand.Intn(1000))
	return utils.Md5(str)
}

func (s *AdminService) InfoByToken(t string) *model.Admin {
	a := &model.Admin{}
	global.DB.Scopes(CommonEnable()).Where("token = ?", t).First(a)
	return a
}

func (s *AdminService) CheckToken(t string) (*model.Admin, bool) {
	a := s.InfoByToken(t)
	if a.Id > 0 {
		exp := a.TokenExpireTime
		if time.Now().Sub(exp) < 0 {
			return a, true
		}
	}
	return a, false
}

func (s *AdminService) Update(a *model.Admin, f *request.AdminForm) error {
	v := map[string]interface{}{
		"role_id":  f.RoleId,
		"username": f.Username,
		"nickname": f.Nickname,
		"status":   f.Status,
	}
	err := global.DB.Model(a).Updates(v).Error
	return err
}

func (s *AdminService) Create(f *request.AdminForm) (*model.Admin, error) {
	admin := &model.Admin{
		RoleId:   f.RoleId,
		Username: f.Username,
		Password: utils.Md5(f.Password),
		Nickname: f.Nickname,
		Status:   f.Status,
	}
	err := global.DB.Create(admin).Error
	return admin, err
}

func (s *AdminService) Delete(a *model.Admin) error {
	return global.DB.Delete(a).Error
}
