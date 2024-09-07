package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"practise-sample-store/constant"
	"practise-sample-store/entity"
	"practise-sample-store/service"
	"practise-sample-store/util"
)

func UserLogin(ctx *gin.Context) {
	user := &entity.User{}
	if err := ctx.ShouldBindJSON(user); err == nil {
		if validate.IsEmpty(user.Username) || validate.IsEmpty(user.Password) {
			ctx.JSON(200, constant.BuildFail("用户名不能为空 或密码不能为空"))
			return
		}
		if ok, _ := service.UserLogin(user.Username, user.Password); ok {
			userToken := util.GenerateTokenWithTimestamp(user.Username)
			service.UserTokenCache.Set(user.Username, userToken)
			ctx.JSON(200, constant.BuildSuccess(userToken))
			return
		}
	}
	ctx.JSON(200, constant.BuildFail("登陆失败"))
}

func UserInfo(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")

	user := &entity.User{}
	if err := ctx.ShouldBindJSON(user); err == nil {
		if validate.IsEmpty(user.Username) {
			ctx.JSON(200, constant.BuildFail("用户名不能为空"))
			return
		}
		if tokenCache, _ := service.UserTokenCache.Get(user.Username); tokenCache.(string) != token {
			ctx.JSON(200, constant.BuildFail("token 校验不正确"))
			return
		}

		if userInfo, err := service.GetUserInfo(user.Username); err == nil {
			ctx.JSON(200, constant.BuildSuccess(userInfo))
			return
		}
		ctx.JSON(200, constant.BuildFail("用户不存在"))

	}
	ctx.JSON(200, constant.BuildFail("参数不正确"))

}
