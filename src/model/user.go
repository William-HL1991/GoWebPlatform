package model

import (
	"errors"
)

type User struct {
	Id    int64
	Name  string  `xorm:"varchar(25) notnull unique 'usr_name'"`

}

// 创建用户
func newUser(name string) error {
	// 对未存在的记录进行插入
	_, err := x.Insert(&User{Name:name})
	return err
}

// 获取用户信息
func getUser(id int64)(*User, error) {
	u := &User{}
	// 直接操作 ID
	has, err := x.Id(id).Get(u)
	// 兼容对象的出错
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("User don`t exist")
	}
	return u, nil

}

// 删除用户
func deleteUser(id int64) error {
	// 通过 delete 方法删除记录
	_, err := x.Delete(&User{Id: id})
	return err
}