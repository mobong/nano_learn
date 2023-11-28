package logic

import (
	"fmt"
	"github.com/lonng/nano/session"
	"nano_learn/db/model"
	"nano_learn/proto"
	"time"
)

// 登录
func (rm *RoleManager) Login(s *session.Session, req *proto.LoginTos) error {
	uid := req.Uid
	s.Bind(int64(uid))
	user, ok, err := model.QueryUser(uid)
	if ok {
		user = &model.User{
			UId:      uid,
			Username: req.Username,
			Status:   1,
			Ctime:    time.Now(),
			IsOnline: true,
		}
		model.InsertUser(user)
	} else if err != nil {
		return err
	}

	res := &proto.LoginToc{
		Code:     0,
		Uid:      uint64(s.UID()),
		Username: req.Username,
	}
	return s.Response(res)
}

// 创建role
func (rm *RoleManager) CreateRole(s *session.Session, req *proto.CreateRoleTos) error {
	uid := req.Uid
	role, ok, err := model.QueryRole(uid)
	if ok {
		if role != nil {
			role.LoginTime = time.Now()
		} else {
			role = newRole(s, req.Uid, req.Name, uint8(req.Sex))
		}
		rm.setRole(uid, role)
	} else if err != nil {
		return err
	}
	// 添加到广播频道
	rm.Group.Add(s)

	res := &proto.CreateRoleToc{
		Code: 0,
		Rid:  role.RId,
		Uid:  uint64(s.UID()),
		Name: req.Name,
	}
	return s.Response(res)
}

func (rm *RoleManager) role(uid uint64) (*model.Role, bool) {
	role, ok := rm.Roles[uid]

	return role, ok
}

func (rm *RoleManager) setRole(uid uint64, role *model.Role) {
	if _, ok := rm.Roles[uid]; ok {
		fmt.Println("已经存在")
	}
	rm.Roles[uid] = role
}
