package router

import (
	"github.com/gin-gonic/gin"
	"practise-sample-store/controller"
	"practise-sample-store/service"
)

const LOGIN = "login"

const REGISTER = "register"

const USERINFO = "userInfo"

func Router() *gin.Engine {
	engine := gin.Default()

	// login
	engine.POST(LOGIN, controller.UserLoginEndpoint)
	engine.POST(REGISTER, controller.UserRegisterEndpoint)

	// user group
	userGroup := engine.Group("/user")
	userGroup.Use(service.AuthHandler())
	{
		userGroup.GET(USERINFO, controller.UserInfoEndpoint)
	}
	return engine
}
