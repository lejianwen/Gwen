package model

import (
	"database/sql/driver"
	"time"
)

type StatusCode int

const (
	COMMON_STATUS_ENABLE   StatusCode = 1 //通用状态 启用
	COMMON_STATUS_DISABLED StatusCode = 2 //通用状态 禁用
)

type Model struct {
	Id        uint     `gorm:"primaryKey" json:"id"`
	CreatedAt AutoTime `json:"created_at"`
	UpdatedAt AutoTime `json:"updated_at"`
}

// Pagination
type Pagination struct {
	Page      int64 `form:"page" json:"page"`
	TotalSize int64 `form:"total_size" json:"total_size"`
	PageSize  int64 `form:"page_size" json:"page_size"`
}

// AutoTime 自定义时间格式
type AutoTime time.Time

func (mt AutoTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	t := time.Time(mt)
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t, nil
}

func (mt AutoTime) MarshalJSON() ([]byte, error) {
	//b := make([]byte, 0, len("2006-01-02 15:04:05")+2)
	b := time.Time(mt).AppendFormat([]byte{}, "\"2006-01-02 15:04:05\"")
	return b, nil
}
