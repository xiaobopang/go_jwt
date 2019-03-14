package libs
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "13306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	err := orm.RegisterDataBase("default", "mysql", dsn, 30, 30) //注册默认数据库 30连接数
	if err != nil {
		beego.Error("数据库连接错误：", err)
		os.Exit(2)
		return
	}
	//orm.RegisterModel(new(models.User))   //注册model实体
	// err = orm.RunSyncdb("default", false, true) //自动同步表结构
	// if err != nil {
	// 	fmt.Println("自动建表失败：", err)
	// }
	// orm.RunCommand()
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
