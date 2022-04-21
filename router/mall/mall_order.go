package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallOrderRouter struct {
}

func (m *MallOrderRouter) InitMallOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1")
	var mallOrderRouterApi = v1.ApiGroupApp.MallApiGroup.MallOrderApi
	{
		mallOrderRouter.GET("/paySuccess", mallOrderRouterApi.PaySuccess)             //模拟支付成功回调的接口
		mallOrderRouter.PUT("/order/:orderNo/finish", mallOrderRouterApi.FinishOrder) //确认收货接口
		mallOrderRouter.PUT("/order/:orderNo/cancel", mallOrderRouterApi.CancelOrder) //取消订单接口

	}
}
