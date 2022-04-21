package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
	"strconv"
)

type MallOrderApi struct {
}

func (m *MallOrderApi) PaySuccess(c *gin.Context) {
	var req mallReq.PaySuccessParams
	_ = c.ShouldBindQuery(&req)
	if err := mallOrderService.PaySuccess(req.OrderNo, req.PayType); err != nil {
		global.GVA_LOG.Error("订单支付失败", zap.Error(err))
		response.FailWithMessage("订单支付失败:"+err.Error(), c)
	}
	response.OkWithMessage("订单支付成功", c)
}

func (m *MallOrderApi) FinishOrder(c *gin.Context) {
	var orderNo string
	_ = c.ShouldBindQuery("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.FinishOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("订单签收失败", zap.Error(err))
		response.FailWithMessage("订单签收失败:"+err.Error(), c)
	}
	response.OkWithMessage("订单签收成功", c)

}

func (m *MallOrderApi) CancelOrder(c *gin.Context) {
	var orderNo string
	_ = c.ShouldBindQuery("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.CancelOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("订单签收失败", zap.Error(err))
		response.FailWithMessage("订单签收失败:"+err.Error(), c)
	}
	response.OkWithMessage("订单签收成功", c)

}

func (m *MallOrderApi) OrderList(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.OrderSearchParams
	_ = c.ShouldBindQuery(&req)
	status, _ := strconv.Atoi(req.Status)
	if !utils.NumsInList(status, []int{0, 1, 2, 3, 4}) {
		response.FailWithMessage("非法搜索参数", c)
	}
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if err, list, total := mallOrderService.MallOrderListBySearch(token, req); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   req.PageNumber,
			PageSize:   5,
		}, "获取成功", c)
	}

}
