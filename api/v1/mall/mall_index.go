package mall

import (
	"github.com/gin-gonic/gin"
	"main.go/model/common/response"
)

type MallIndexApi struct {
}

// MallIndexInfo 加载首页信息
func (m *MallIndexApi) MallIndexInfo(c *gin.Context) {
	response.OkWithMessage("hello mall", c)

}
