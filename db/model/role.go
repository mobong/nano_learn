package model

import "nano_learn/db"

func QueryRole(uid uint64) (*Role, bool, error) {
	role := &Role{}
	has, err := db.Database.Table(role).Where("uid=?", uid).Get(role)
	return role, has, err

}

// UpdateRole
func UpdateRole(r *Role) error {
	if r == nil {
		return nil
	}
	_, err := db.Database.Table(r).Where("uid=?", r.UId).AllCols().Update(r)
	return err
}

// InsertRole
func InsertRole(r *Role) error {
	if r == nil {
		return nil
	}
	_, err := db.Database.Table(r).Insert(r)
	return err
}

// DeleteRole
func DeleteRole(rid int64) error {
	r := &Role{}
	_, err := db.Database.Table(r).Where("rid=?", rid).Delete(r)
	return err
}
