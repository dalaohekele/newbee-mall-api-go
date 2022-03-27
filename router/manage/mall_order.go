package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallOrderRouter struct {
}

// InitMallOrderRouter 初始化 MallOrder 路由信息
func (s *MallOrderRouter) InitMallOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1")
	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.MallOrderApi
	{
		mallOrderRouter.PUT("orders/checkDone", mallOrderApi.CheckDoneOrder) // 发货
		mallOrderRouter.PUT("orders/checkOut", mallOrderApi.CheckOutOrder)   // 出库
		mallOrderRouter.PUT("orders/close", mallOrderApi.CloseOrder)         // 出库
		mallOrderRouter.GET("orders/:orderId", mallOrderApi.FindMallOrder)   // 根据ID获取MallOrder
		mallOrderRouter.GET("orders", mallOrderApi.GetMallOrderList)         // 获取MallOrder列表
	}
}
