package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallGoodsInfoRouter struct {
}

// InitMallGoodsInfoRouter 初始化 MallGoodsInfo 路由信息
func (s *MallGoodsInfoRouter) InitMallGoodsInfoRouter(Router *gin.RouterGroup) {
	mallGoodsInfoRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.ManageApiGroup.MallGoodsInfoApi
	{
		mallGoodsInfoRouter.POST("goods", mallGoodsInfoApi.CreateMallGoodsInfo)                    // 新建MallGoodsInfo
		mallGoodsInfoRouter.DELETE("deleteMallGoodsInfo", mallGoodsInfoApi.DeleteMallGoodsInfo)    // 删除MallGoodsInfo
		mallGoodsInfoRouter.PUT("goods/status/:status", mallGoodsInfoApi.ChangeMallGoodsInfoByIds) // 上下架
		mallGoodsInfoRouter.PUT("goods", mallGoodsInfoApi.UpdateMallGoodsInfo)                     // 更新MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/:id", mallGoodsInfoApi.FindMallGoodsInfo)                   // 根据ID获取MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/list", mallGoodsInfoApi.GetMallGoodsInfoList)               // 获取MallGoodsInfo列表
	}
}
