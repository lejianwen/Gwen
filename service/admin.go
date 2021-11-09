package service

import (
	"Gwen/global"
	"Gwen/model"
	"Gwen/utils"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type AdminService struct {
}

func (s *AdminService) Info(id int) *model.Admin {
	a := &model.Admin{}
	global.DB.Scopes(CommonEnable()).Where("id = ?", id).First(a)
	return a
}

func (s *AdminService) List() (res *model.AdminListRes) {
	res = &model.AdminListRes{}
	global.DB.Scopes(Paginate(1, 10)).Preload("Role").Find(&res.Admins)
	global.DB.Model(&model.Admin{}).Count(&res.TotalSize)
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
