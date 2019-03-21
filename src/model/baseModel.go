package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

var x *xorm.Engine

func init()  {
	// 创建 ORM 引擎与数据库
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "", "localhost:3306", "LRP")
	var err error
	x, err = xorm.NewEngine("mysql", params)
	if err != nil {
		log.Fatal("Fatal to creat engin: %v\n", err)
	}

	// 添加统一前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "LRP_")
	x.SetTableMapper(tbMapper)
	// defer x.Close() 这个地方不能每次都关闭，要等所有任务完成才能关闭

	// 同步结构体与数据表 User
	if err = x.Sync(new(User)); err != nil {
		log.Fatal("Fail to sync database: %v\n", err)
	}

	x.ShowSQL(true)
	x.ShowExecTime(true)

}
