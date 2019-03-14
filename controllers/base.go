package controllers
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import (
	. "go_jwt/libs"
	"time"

	"github.com/astaxie/beego"
)

const TimeFormart = "2006-01-02 15:04:05"

// Operations about Main
type MainController struct {
	beego.Controller
}

//map[string]interface{}{"code": 400, "msg": "no user exists!", "data": nil}
type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	TimeStamp int64       `json:"timestamp"`
}

func OutResponse(code int, data interface{}, msg string) Response {
	if msg == "" {
		msg = GetMsg(code)
	}
	Resp := Response{
		Code:      code,
		Msg:       msg,
		Data:      data,
		TimeStamp: time.Now().Unix(), //time.Now().Format("2006-01-02 15:04:05")
	}
	return Resp
}

func (m *MainController) Welcome() {
	// m.Data["Website"] = "www.unclepang.com"
	// m.Data["Email"] = "10846295@qq.com"
	// m.TplName = "index.tpl"
	// m.Render()
	m.Data["json"] = OutResponse(1, nil, "")
	m.ServeJSON()
}
