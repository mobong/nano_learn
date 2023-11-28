package model

import (
	"nano_learn/db"
)

func QueryUser(uid uint64) (*User, bool, error) {
	user := &User{}
	has, err := db.Database.Table(user).Where("uid=?", uid).Get(user)
	return user, has, err

}

// UpdateUser
func UpdateUser(u *User) error {
	if u == nil {
		return nil
	}
	_, err := db.Database.Table(u).Where("uid=?", u.UId).AllCols().Update(u)
	return err
}

// InsertUser
func InsertUser(u *User) error {
	if u == nil {
		return nil
	}
	_, err := db.Database.Table(u).Insert(u)
	return err
}

// DeleteUser
func DeleteUser(uid int64) error {
	u := &User{}
	_, err := db.Database.Table(u).Where("uid=?", uid).Delete(u)
	return err
}
