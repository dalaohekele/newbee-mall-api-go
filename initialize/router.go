package initialize

import (
	"github.com/gin-gonic/gin"
	"main.go/global"
	"main.go/router"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.GVA_LOG.Info("use middleware logger")
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	//商城后管路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage-api")
	{
		//商城后管路由初始化
		manageRouter.InitManageRouter(ManageGroup)
	}
	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	{
		// 商城前端路由
		mallRouter.InitMallRouter(MallGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
