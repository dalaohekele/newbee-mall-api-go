package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type ManageRouter struct {
}

func (r *ManageRouter) InitManageRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1")
	var manageApi = v1.ApiGroupApp.ManageApiGroup.ManageApi
	{
		mallAdminUserRouter.GET("manage", manageApi.TestApi)
	}

}
