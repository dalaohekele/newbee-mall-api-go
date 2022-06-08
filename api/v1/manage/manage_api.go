package manage

import (
	"github.com/gin-gonic/gin"
	"main.go/model/common/response"
)

type ManageApi struct {
}

// 创建AdminUser
func (m *ManageApi) TestApi(c *gin.Context) {
	response.OkWithMessage("hello manage", c)
}
