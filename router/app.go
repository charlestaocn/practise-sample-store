package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	engine := gin.Default()
	UserRoute(engine)
	UserInfoRoute(engine)
	return engine
}
