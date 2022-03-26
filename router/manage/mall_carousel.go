package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallCarouselRouter struct {
}

// InitMallCarouselRouter 初始化 轮播图模块 路由信息
func (s *MallCarouselRouter) InitMallCarouselRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallCarouselApi = v1.ApiGroupApp.ManageApiGroup.MallCarouselApi
	{
		mallCarouselRouter.POST("carousels", mallCarouselApi.CreateMallCarousel)   // 新建MallCarousel
		mallCarouselRouter.DELETE("carousels", mallCarouselApi.DeleteMallCarousel) // 删除MallCarousel
		mallCarouselRouter.PUT("carousels", mallCarouselApi.UpdateMallCarousel)    // 更新MallCarousel
		mallCarouselRouter.GET("carousels/:id", mallCarouselApi.FindMallCarousel)  // 根据ID获取轮播图
		mallCarouselRouter.GET("carousels", mallCarouselApi.GetMallCarouselList)   // 获取轮播图列表
	}
}
