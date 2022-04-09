package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallGoodsInfoApi struct {
}

func (m *MallGoodsInfoApi) GoodsSearch(c *gin.Context) {
	var req mallReq.GoodsSearchParams
	_ = c.ShouldBindQuery(&req)
	if req.Keyword == "" {
		response.FailWithMessage("非法搜索参数", c)
	}
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if err, list, total := mallGoodsInfoService.MallGoodsListBySearch(req); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   req.PageNumber,
			PageSize:   10,
		}, "获取成功", c)
	}

}

func (m *MallGoodsInfoApi) GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}
	response.OkWithData(goodsInfo, c)
}
