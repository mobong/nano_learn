package logic

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"nano_learn/db/model"
)

// logic manager
type RoleManager struct {
	component.Base
	*nano.Group
	Roles map[uint64]*model.Role
}

func NewManager() *RoleManager {
	return &RoleManager{
		Roles: map[uint64]*model.Role{},
	}
}
