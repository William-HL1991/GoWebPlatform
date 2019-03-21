package model

import (
	"errors"
)

type User struct {
	Id        int64
	Name      string  `xorm:"varchar(25) notnull unique 'usrname'"`
	Password  string  `xorm:"varchar(16) notnull unique 'password'"`

}

// 创建用户
func NewUser(name string, passwd string) error {
	// 对未存在的记录进行插入
	_, err := x.Insert(&User{Name:name, Password:passwd})
	return err
}

// 获取用户信息
func GetUserById(id int64)(*User, error) {
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

//通过用户名查询用户
func GetUserByName(name string)(bool, error) {
	// 查询用户名
	has, err := x.Exist(&User{Name:name})
	return has, err
}

// 删除用户
func DeleteUser(id int64) error {
	// 通过 delete 方法删除记录
	_, err := x.Delete(&User{Id: id})
	return err
}
