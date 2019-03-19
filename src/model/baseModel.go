package model

import (
	"github.com/go-xorm/xorm"
	"log"
)

var x *xorm.Engine

func init()  {
	// 创建 ORM 引擎与数据库
	var err error
	x ,err = xorm.NewEngine("mysql", "root:11111@/sys?charset=utf-8")
	if err != nil {
		log.Fatal("Fatal to creat engin: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(User)); err != nil {
		log.Fatal("Fail to sync database: %v\n", err)
	}
}

