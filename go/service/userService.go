package service

import (
	"checkin/middleware"
	"checkin/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserService struct{}

func (UserService) Login(ctx *gin.Context) {
	request := make(map[string]interface{})
	err := ctx.ShouldBind(&request)
	if err != nil {
		return
	}
	user := model.User{}
	result := model.DB.Where("id = ? and password = ?", request["id"], request["password"]).Find(&user)
	if result.RowsAffected != 1 {
		ctx.JSON(200, new(model.Result).Error("登录失败"))
		return
	}
	claims := jwt.MapClaims{
		"id":  "Callay",
		"pwd": "020804",
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	token, err := middleware.GenerateJwtToken(&claims)
	if err != nil {
		ctx.JSON(200, new(model.Result).Error("jwt令牌生成错误"))
		return
	}
	ctx.JSON(200, new(model.Result).OKSetMsgData("登录成功", token))
}

func (UserService) GetUserInfo(ctx *gin.Context) {
	id, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(200, new(model.Result).Error("用户信息获取失败"))
		return
	}
	user := model.User{Id: id.(string)}
	model.DB.First(&user)
	ctx.JSON(200, new(model.Result).OKSetData(user))
}
