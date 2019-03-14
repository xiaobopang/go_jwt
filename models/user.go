package models
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import (
	"go_jwt/libs"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int    `json:id`
	Name      string `json:name`
	Email     string `json:email`
	Age       int    `json:age`
	CreatedAt int64  `json:created_at`
	UpdatedAt int64  `json:updated_at`
}

func (m *User) TableName() string {
	return libs.TableName("user")
}

func init() {
	orm.RegisterModel(new(User))
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (u *User, err error) {
	o := orm.NewOrm()
	//o.Using("default2") // 默认使用 default，你可以指定为其他数据库
	u = &User{Id: id}
	if err = o.Read(u); err == nil {
		return u, nil
	}
	return nil, err
}

func AddUser(u *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(u)
	if err = o.Read(u); err == nil {
		return id, nil
	}
	return id, err
}
