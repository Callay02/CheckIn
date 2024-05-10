package service

import (
	"checkin/middleware"
	"checkin/model"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	id, _ := ctx.Get("uid")
	user := model.User{Id: id.(string)}
	model.DB.First(&user)
	ctx.JSON(200, new(model.Result).OKSetData(user))
}

func (UserService) UpdateUserInfo(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	user := model.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(200, new(model.Result).Error("用户信息获取错误"))
		return
	}
	fmt.Println(user)
	user.Id = id.(string)
	user.UpdateTime = strconv.FormatInt(time.Now().Unix(), 10)
	err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		ctx.JSON(200, new(model.Result).Error("用户信息更新失败"))
		return
	}
	ctx.JSON(200, new(model.Result).OK("用户信息更新成功"))
}
