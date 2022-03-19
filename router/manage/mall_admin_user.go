package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallAdminUserRouter struct {
}

// InitMallAdminUserRouter 初始化 MallAdminUser 路由信息
func (s *MallAdminUserRouter) InitMallAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth()).Use(middleware.Cors())
	mallAdminUserWithoutRouter := Router.Group("v1")
	var mallAdminUserApi = v1.ApiGroupApp.ManageApiGroup.MallAdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateMallAdminUser) // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateMallAdminUserName)   // 更新MallAdminUser
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateMallAdminUserPassword)
		mallAdminUserRouter.GET("users", mallAdminUserApi.UserList)
		mallAdminUserRouter.PUT("users/:lockStatus", mallAdminUserApi.LockUser)
		mallAdminUserRouter.GET("adminUser/profile", mallAdminUserApi.AdminUserProfile) // 根据ID获取 admin详情
		mallAdminUserWithoutRouter.DELETE("logout", mallAdminUserApi.AdminLogout)
	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin) //管理员登陆
	}
}
