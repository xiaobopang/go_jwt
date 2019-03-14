package controllers

import (
	"encoding/json"
	"fmt"
	"go_jwt/libs"
	"go_jwt/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) URLMapping() {
	u.Mapping("Get", u.Get)
	u.Mapping("Post", u.Post)
	u.Mapping("Redis", u.Redis)
	u.Mapping("Login", u.Login)
	u.Mapping("ParseJwt", u.ParseJwt)
}

// @router /:uid [get]
func (u *UserController) Get() {
	//u.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", u.Ctx.Request.Header.Get("Origin"))
	userId, _ := u.GetInt(":uid")
	if userId == 0 {
		u.Data["json"] = OutResponse(404, nil, "")
		u.ServeJSON()
		return
	} else {
		res, err := models.GetUserById(userId)
		if err != nil {
			u.Data["json"] = OutResponse(400, nil, "")
			u.ServeJSON()
			return
		}
		//name,value,time, path,domain, secure and httponly.
		u.Ctx.SetCookie("name", res.Name, 100, "/", "", false, true)   // 设置cookie
		u.Ctx.SetCookie("email", res.Email, 100, "/", "", false, true) // 设置cookie

		//u.Ctx.Output.SetStatus(400)
		u.Data["json"] = OutResponse(200, res, "")
		u.ServeJSON()
		return
	}
}

// @router /create [post]
func (u *UserController) Post() {
	var NewUser models.User

	NewUser.Name = u.GetString("Name")
	NewUser.Age, _ = u.GetInt("Age")
	NewUser.Email = u.GetString("Email")
	NewUser.CreatedAt = time.Now().Unix()
	NewUser.UpdatedAt = time.Now().Unix()
	newId, err := models.AddUser(&NewUser)

	if err != nil {
		u.Data["json"] = OutResponse(400, nil, "")
		u.ServeJSON()
		return
	}
	u.Data["json"] = OutResponse(200, newId, "")
	u.ServeJSON()
	return
}

// @router /redis/:redisKey [get]
func (u *UserController) Redis() {
	var user models.User
	redisKey := u.GetString(":redisKey")
	userInfo, err := libs.GetKey(redisKey)
	json.Unmarshal([]byte(userInfo), &user)
	// fmt.Println("token 666===>", libs.GenerateToken(user.Id))
	// fmt.Println(libs.ParseToken(libs.GenerateToken(user.Id)))
	if err != nil {
		user, err := models.GetUserById(1)
		userJson, _ := json.Marshal(&user)
		libs.SetKey(redisKey, userJson, 3600)
		if err != nil {
			fmt.Println(err)
		}
		u.Data["json"] = OutResponse(200, &user, "")
		u.ServeJSON()
		return
	} else {
		user, err := models.GetUserById(1)
		if err != nil {
			fmt.Println(err)
		}
		u.Data["json"] = OutResponse(200, &user, "")
		u.ServeJSON()
		return

	}
	u.Data["json"] = OutResponse(200, &user, "")
	u.ServeJSON()
}

// @Title login
// @Description Logs out current logged in user session
// @Success 200 {string} login success
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	userId, _ := u.GetInt("id")

	if username != "" && password != "" {
		//fmt.Println("加密后", libs.Sha256(libs.Md5(password)+beego.AppConfig.String("jwt::salt")))
		token := libs.GenerateToken(userId, u.Ctx.Input.Domain())
		u.Data["json"] = OutResponse(200, map[string]string{"token": token}, "")
		u.ServeJSON()
	} else {
		u.Data["json"] = OutResponse(400, nil, "用户名或密码不能为空")
		u.ServeJSON()
	}
}

// @Title parseJwt
// @Description Logs out current logged in user session
// @Success 200 {string} login success
// @router /parse_jwt [get]
func (u *UserController) ParseJwt() {
	tokenString := u.Ctx.Input.Header("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(beego.AppConfig.String("jwt::token")), nil
	})
	fmt.Println(err)
	userId := libs.GetIdFromClaims("user_id", token.Claims)
	u.Data["json"] = OutResponse(200, userId, "")
	u.ServeJSON()
}
