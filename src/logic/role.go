package logic

import (
	"github.com/lonng/nano/session"
	"nano_learn/db/model"
	"time"
)

func newRole(s *session.Session, uid uint64, name string, sex uint8) *model.Role {
	p := &model.Role{
		UId:         uid,
		Sex:         sex,
		Name:        name,
		LoginTime:   time.Now(),
		CreatedTime: time.Now(),
		Session:     s,
	}
	return p
}
