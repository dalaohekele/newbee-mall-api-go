package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallRouter struct {
}

func (m *MallRouter) InitMallRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1")
	var mallCarouselApi = v1.ApiGroupApp.MallApiGroup.MallIndexApi
	{
		mallCarouselRouter.GET("mall", mallCarouselApi.MallIndexInfo) // 获取首页数据
	}
}
