package middlewares
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import (
	"encoding/json"
	"fmt"
	"go_jwt/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

func init() {
}

func ValidateToken(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	var uri string = ctx.Input.URI()
	if uri == "/v1/user/login" {
		return
	}

	if ctx.Input.Header("Authorization") == "" {
		ctx.Output.SetStatus(401)
		resBody, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求,token不合法"))
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	var tokenString string = ctx.Input.Header("Authorization")
	// validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(beego.AppConfig.String("jwt::token")), nil
	})
	// branch out into the possible error from signing
	switch err.(type) {

	case nil: // no error
		if !token.Valid { // but may still be invalid
			ctx.Output.SetStatus(401)
			resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
			ctx.Output.Body(resBytes)
			if err != nil {
				panic(err)
			}
		}
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			ctx.Output.SetStatus(401)
			resBody, err := json.Marshal(controllers.OutResponse(401, nil, "登录已过期，请重新登录"))
			ctx.Output.Body(resBody)
			if err != nil {
				panic(err)
			}
		default:
			ctx.Output.SetStatus(401)
			resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
			ctx.Output.Body(resBytes)
			if err != nil {
				panic(err)
			}
		}
	default: // something else went wrong
		ctx.Output.SetStatus(401)
		resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
		ctx.Output.Body(resBytes)
		if err != nil {
			panic(err)
		}
	}
	if err == nil && token.Valid {
		return
	}
}
