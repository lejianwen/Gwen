package model

import (
	_ "gorm.io/driver/mysql"
	"time"
)

type StatusCode int

const (
	COMMON_STATUS_ENABLE   StatusCode = 1 //通用状态 启用
	COMMON_STATUS_DISABLED StatusCode = 2 //通用状态 禁用
)

type Model struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Pagination
type Pagination struct {
	PageNum   int64 `form:"page_num" json:"page_num"`
	TotalSize int64 `form:"total_size" json:"total_size"`
	PageSize  int64 `form:"page_size" json:"page_size"`
}
