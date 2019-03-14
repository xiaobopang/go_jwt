package libs
/*
 * Created Date: Wednesday March 13th 2019
 * Author: Pangxiaobo
 * Last Modified: Wednesday March 13th 2019 2:52:45 pm
 * Modified By: the developer formerly known as Pangxiaobo at <10846295@qq.com>
 * Copyright (c) 2019 Pangxiaobo
 */

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/dgrijalva/jwt-go"
)

func init() {}

// iss: jwt签发者
// sub: jwt所面向的用户
// aud: 接收jwt的一方
// exp: jwt的过期时间，这个过期时间必须要大于签发时间
// nbf: 定义在什么时间之前，该jwt都是不可用的.
// iat: jwt的签发时间
// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
func GenerateToken(userId int, domain string) string {
	var tokenExpire int64
	// current timestamp
	currentTimestamp := time.Now().UTC().Unix()
	tokenExpire, err := beego.AppConfig.Int64("jwt::token_expire")
	if err != nil {
		tokenExpire = 600
	}
	// md5 of sub & iat
	h := md5.New()
	io.WriteString(h, strconv.Itoa(userId))
	io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iat": currentTimestamp,
		"exp": currentTimestamp + tokenExpire,
		"nbf": currentTimestamp,
		"iss": domain,
		"jti": h.Sum(nil),
		"user_id":userId,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("jwt::token")))

	if err != nil {
		log.Fatal(err)
	}

	return (tokenString)
}

// (获得payload的信息)从token对象里获得参数(key)对应的值
func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
