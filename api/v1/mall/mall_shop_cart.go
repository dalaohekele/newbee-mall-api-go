package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallShopCartApi struct {
}

func (m *MallShopCartApi) CartItemList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, shopCartItem := mallShopCartService.GetMyShoppingCartItems(token); err != nil {
		global.GVA_LOG.Error("获取购物车失败", zap.Error(err))
		response.FailWithMessage("获取购物车失败:"+err.Error(), c)
	} else {
		response.OkWithData(shopCartItem, c)
	}
}

func (m *MallShopCartApi) CartItemListByPage(c *gin.Context) {
	token := c.GetHeader("token")
	var pageInfo mallReq.MallShopCartSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallShopCartService.GetShopCartListByPage(token, pageInfo); err != nil {
		global.GVA_LOG.Error("获取购物车失败", zap.Error(err))
		response.FailWithMessage("获取购物车失败:"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (m *MallShopCartApi) SaveMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.SaveCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.SaveMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("添加购物车失败", zap.Error(err))
		response.FailWithMessage("添加购物车失败:"+err.Error(), c)
	}
	response.OkWithMessage("添加购物车成功", c)
}

func (m *MallShopCartApi) UpdateMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.UpdateCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.UpdateMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("修改购物车失败:"+err.Error(), c)
	}
	response.OkWithMessage("修改购物车成功", c)
}

func (m *MallShopCartApi) DelMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(token, id); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("修改购物车失败:"+err.Error(), c)
	}
	response.OkWithMessage("修改购物车成功", c)
}

// todo 这里的传参貌似有异常,需要调试
func (m *MallShopCartApi) ToSettle(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if len(IDS.Ids) < 1 {
		response.FailWithMessage("参数异常", c)
	}
	token := c.GetHeader("token")
	if err, cartItemRes := mallShopCartService.GetCartItemsForSettle(token, IDS.Ids); err != nil {
		global.GVA_LOG.Error("获取购物明细异常：", zap.Error(err))
		response.FailWithMessage("获取购物明细异常:"+err.Error(), c)
	} else {
		response.OkWithData(cartItemRes, c)
	}

}
