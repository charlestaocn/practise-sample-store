package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"practise-sample-store/constant"
	"practise-sample-store/entity"
	"practise-sample-store/service"
	"practise-sample-store/util"
)

func UserLoginEndpoint(ctx *gin.Context) {
	user := &entity.User{}
	if err := ctx.ShouldBindJSON(user); err == nil {
		if validate.IsEmpty(user.Username) || validate.IsEmpty(user.Password) {
			ctx.JSON(200, constant.BuildFail("用户名不能为空 或密码不能为空"))
			return
		}
		if ok, _ := service.UserLoginService(user.Username, user.Password); ok {
			userToken := util.GenerateTokenWithTimestamp(user.Username)
			service.UserTokenCache.Set(user.Username, userToken)
			ctx.JSON(200, constant.BuildSuccess(userToken))
			return
		}
	}
	ctx.JSON(200, constant.BuildFail("登陆失败"))
}

func UserInfoEndpoint(ctx *gin.Context) {
	user := &entity.User{}
	if err := ctx.ShouldBindJSON(user); err == nil {
		if userInfo, err := service.GetUserInfoService(user.Username); err == nil {
			ctx.JSON(200, constant.BuildSuccess(userInfo))
			return
		}
		ctx.JSON(200, constant.BuildFail("用户不存在"))

	}
	ctx.JSON(200, constant.BuildFail("参数不正确"))

}

func UserRegisterEndpoint(ctx *gin.Context) {
	user := &entity.User{}
	if err := ctx.ShouldBindJSON(user); err == nil {
		registerRes, msg := service.RegisterUserService(user)
		if registerRes {
			ctx.JSON(200, constant.BuildSuccess(msg))
		} else {
			ctx.JSON(200, constant.BuildFail(msg))
		}
		return
	} else {
		ctx.JSON(200, constant.BuildFail("参数不正确"))
	}
}
