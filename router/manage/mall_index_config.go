package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallIndexConfigRouter struct {
}

// InitMallIndexConfigRouter 初始化 MallIndexConfig 路由信息
func (s *MallIndexConfigRouter) InitMallIndexConfigRouter(Router *gin.RouterGroup) {
	mallIndexConfigRouter := Router.Group("v1")
	var mallIndexConfigApi = v1.ApiGroupApp.ManageApiGroup.MallIndexConfigApi
	{
		mallIndexConfigRouter.POST("indexConfigs", mallIndexConfigApi.CreateMallIndexConfig)        // 新建MallIndexConfig
		mallIndexConfigRouter.POST("indexConfigs/delete", mallIndexConfigApi.DeleteMallIndexConfig) // 删除MallIndexConfig
		mallIndexConfigRouter.PUT("indexConfigs", mallIndexConfigApi.UpdateMallIndexConfig)         // 更新MallIndexConfig
		mallIndexConfigRouter.GET("indexConfigs/:id", mallIndexConfigApi.FindMallIndexConfig)       // 根据ID获取MallIndexConfig
		mallIndexConfigRouter.GET("indexConfigs", mallIndexConfigApi.GetMallIndexConfigList)        // 获取MallIndexConfig列表
	}
}
