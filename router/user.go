package router

import (
	"github.com/gin-gonic/gin"
	"practise-sample-store/controller"
)

func UserRoute(router *gin.Engine) *gin.Engine {
	router.POST(LOGIN, controller.UserLogin)
	return router
}

func UserInfoRoute(router *gin.Engine) *gin.Engine {
	router.GET(USERINFO, controller.UserInfo)
	return router
}
