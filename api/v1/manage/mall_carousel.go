package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
)

type MallCarouselApi struct {
}

// CreateMallCarousel 创建MallCarousel
func (m *MallCarouselApi) CreateMallCarousel(c *gin.Context) {
	var req manageReq.MallCarouselAddParam
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.CarouselAddParamVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mallCarouselService.CreateMallCarousel(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMallCarousel 删除MallCarousel
func (m *MallCarouselApi) DeleteMallCarousel(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallCarouselService.DeleteMallCarousel(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateMallCarousel 更新MallCarousel
func (m *MallCarouselApi) UpdateMallCarousel(c *gin.Context) {
	var req manageReq.MallCarouselUpdateParam
	_ = c.ShouldBindJSON(&req)
	//if err := utils.Verify(req, utils.CarouselAddParamVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := mallCarouselService.UpdateMallCarousel(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMallCarousel 用id查询MallCarousel
func (m *MallCarouselApi) FindMallCarousel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, mallCarousel := mallCarouselService.GetMallCarousel(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(mallCarousel, c)
	}
}

// GetMallCarouselList 分页获取MallCarousel列表
func (m *MallCarouselApi) GetMallCarouselList(c *gin.Context) {
	var pageInfo manageReq.MallCarouselSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallCarouselService.GetMallCarouselInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
