package controllers
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import "github.com/astaxie/beego"

// ErrorController definition.
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = OutResponse(404, nil, "METHOD NOT FOUND")
	c.ServeJSON()
}

func (c *ErrorController) Error401() {
	c.Data["json"] = OutResponse(401, nil, "PERMISSION DENIEND")
	c.ServeJSON()
}

func (c *ErrorController) Error403() {
	c.Data["json"] = OutResponse(403, nil, "FORBBIDEN")
	c.ServeJSON()
}

func (c *ErrorController) Error429() {
	c.Data["json"] = OutResponse(429, nil, "Too Many Requests")
	c.ServeJSON()
}
