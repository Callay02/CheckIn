/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 14:01:33
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 22:05:26
 * @FilePath: \checkIn\go\middleware\jwt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"checkin/config"
	"checkin/model"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Key = config.Application.Jwt.Key

// 生成token
func GenerateJwtToken(claims *jwt.MapClaims) (token string, err error) {
	then := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = then.SignedString([]byte(Key))
	if err != nil {
		return token, err
	}
	return token, nil
}

// 解析token
func ParseJwtToken(token string) (claims jwt.MapClaims, err error) {
	tokn, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	claim, ok := tokn.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析错误")
		return claim, err
	}
	if !tokn.Valid {
		err = errors.New("令牌错误！")
		return claim, err
	}
	return claim, nil
}

func JwtTokenVerification(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Abort()
			ctx.JSON(200, new(model.Result).Error("jwt-token验证错误"))
			return
		}
	}()
	claims, err := ParseJwtToken(ctx.Request.Header.Get("jwt-token"))
	if err != nil {
		ctx.Abort()
		ctx.JSON(200, new(model.Result).Error("jwt-token解析错误"))
		return
	}
	ctx.Set("uid", claims["id"])
}
