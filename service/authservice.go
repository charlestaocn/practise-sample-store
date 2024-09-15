package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"practise-sample-store/constant"
)

func AuthHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Print("auth")
		token := context.GetHeader("Authorization")
		username := context.GetHeader("Username")

		if validate.IsEmpty(username) {
			context.JSON(200, constant.BuildFail("用户名不能为空"))
			context.Abort()
			return
		}

		if tokenCache, ok := UserTokenCache.Get(username); !ok || (ok && tokenCache.(string) != token) {
			context.JSON(200, constant.BuildFail("token 校验不正确"))
			context.Abort()
			return
		}

		context.Next()

	}

}
