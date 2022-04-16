package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		mallUserRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate)    //修改用户信息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)       //获取用户信息
		mallUserRouter.POST("/user/login", mallUserApi.UserLogin)       //登陆
		mallUserRouter.POST("/user/logout", mallUserApi.UserLogout)     //登出
	}

}
