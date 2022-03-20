package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/enum"
	"main.go/model/common/request"
	"main.go/model/common/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	manageRes "main.go/model/manage/response"
	"main.go/utils"
	"strconv"
)

type GoodsCategoryApi struct {
}

// CreateCategory 新建商品分类
func (g *GoodsCategoryApi) CreateCategory(c *gin.Context) {
	var category manageReq.MallGoodsCategoryReq
	_ = c.ShouldBindJSON(&category)
	if err := utils.Verify(category, utils.GoodsCategoryVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsCategoryService.AddCategory(category); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateCategory 修改商品分类信息
func (g *GoodsCategoryApi) UpdateCategory(c *gin.Context) {
	var category manage.MallGoodsCategory
	_ = c.ShouldBindJSON(&category)
	if err := utils.Verify(category, utils.GoodsCategoryVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsCategoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败，存在相同分类", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetCategoryList 获取商品分类
func (g *GoodsCategoryApi) GetCategoryList(c *gin.Context) {
	var pageInfo manageReq.SearchCategoryParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := goodsCategoryService.SelectCategoryPage(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败！", c)
	} else {
		response.OkWithData(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
			TotalPage:  int(total) / pageInfo.PageSize,
		}, c)
	}
}

// GetCategory 通过id获取分类数据
func (g *GoodsCategoryApi) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsCategory := goodsCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(manageRes.GoodsCategoryResponse{GoodsCategory: goodsCategory}, c)
	}
}

// DelCategory 设置分类失效
func (g *GoodsCategoryApi) DelCategory(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err, _ := goodsCategoryService.DeleteGoodsCategoriesByIds(ids); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

}

// ListForSelect 用于三级分类联动效果制作
func (g *GoodsCategoryApi) ListForSelect(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsCategory := goodsCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	level := goodsCategory.CategoryLevel
	if level == enum.LevelThree.Code() ||
		level == enum.Default.Code() {
		response.FailWithMessage("参数异常", c)
	}
	categoryResult := make(map[string]interface{})
	if level == enum.LevelOne.Code() {
		_, levelTwoList := goodsCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelTwo.Code())
		if levelTwoList != nil {
			_, levelThreeList := goodsCategoryService.SelectByLevelAndParentIdsAndNumber(int(levelTwoList[0].CategoryId), enum.LevelThree.Code())
			categoryResult["secondLevelCategories"] = levelTwoList
			categoryResult["thirdLevelCategories"] = levelThreeList
		}
	}
	if level == enum.LevelTwo.Code() {
		_, levelThreeList := goodsCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelThree.Code())
		categoryResult["thirdLevelCategories"] = levelThreeList
	}
	response.OkWithData(categoryResult, c)
}
