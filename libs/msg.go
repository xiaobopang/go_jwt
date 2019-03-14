package libs

/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

const (
	WELCOME            = 1
	SUCCESS            = 200
	ERROR              = 500
	INVALID_PARAMS     = 400
	PERMISSION_DENIEND = 401
	NOT_FOUND          = 404
	FORBIDDEN          = 403
	ERROR_EXIST        = 10001
	ERROR_EXIST_FAIL   = 10002
	ERROR_NOT_EXIST    = 10003
	ERROR_GET_S_FAIL   = 10004
	ERROR_ADD_FAIL     = 10006
	ERROR_EDIT_FAIL    = 10007
	ERROR_DELETE_FAIL  = 10008

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	TOO_MANY_REQUESTS              = 429
)

var MsgFlags = map[int]string{
	WELCOME:            "welcome api v1.0.0",
	SUCCESS:            "success",
	ERROR:              "request error",
	PERMISSION_DENIEND: "没有访问权限",
	INVALID_PARAMS:     "请求参数错误",
	NOT_FOUND:          "资源不存在",
	FORBIDDEN:          "禁止访问",
	ERROR_EXIST:        "对象已存在",
	ERROR_EXIST_FAIL:   "获取失败",
	ERROR_NOT_EXIST:    "对象不存在",
	ERROR_GET_S_FAIL:   "获取所有对象失败",
	ERROR_ADD_FAIL:     "新增对象失败",
	ERROR_EDIT_FAIL:    "修改对象失败",
	ERROR_DELETE_FAIL:  "删除对象失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	TOO_MANY_REQUESTS:              "请不要频繁访问",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
