package model

import (
	"github.com/lonng/nano/session"
	"time"
)

type User struct {
	UId      uint64    `xorm:"uid pk autoincr"`
	Username string    `xorm:"username" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	Status   uint8     `xorm:"status"`
	Ctime    time.Time `xorm:"ctime"`
	IsOnline bool      `xorm:"-"`
}

type Role struct {
	RId         uint64           `xorm:"rid pk autoincr"`
	UId         uint64           `xorm:"uid"`
	Sex         uint8            `xorm:"sex"`
	Name        string           `xorm:"name" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	LoginTime   time.Time        `xorm:"login_time"`
	LogoutTime  time.Time        `xorm:"logout_time"`
	CreatedTime time.Time        `xorm:"created_time"`
	Session     *session.Session `xorm:"-"`
}

func (r *Role) TableName() string {
	return "tb_role"
}

func (u *User) TableName() string {
	return "tb_user_info"
}
