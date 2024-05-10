package router

import (
	"checkin/middleware"
	"checkin/service"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	router := r.Group("/user")
	service := service.UserService{}
	{
		router.POST("/login", service.Login)
	}
	{
		router.Use(middleware.JwtTokenVerification)
		router.GET("/jwt")
		router.GET("/getUserInfo", service.GetUserInfo)
		router.POST("/updateUserInfo", service.UpdateUserInfo)
	}
}
