package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	ID         int `xorm:"pk autoincr"`
	UserName   string
	DepartName string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
	Version    int       `xorm:"version"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName should not null!")
	}
	if &u.CreatedAt == nil {
		t := time.Now()
		u.CreatedAt = t
	}
	return &u
}
