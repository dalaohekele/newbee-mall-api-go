package manage

import (
	"errors"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/enum"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"strconv"
	"time"
)

type MallGoodsInfoService struct {
}

// CreateMallGoodsInfo 创建MallGoodsInfo
func (m *MallGoodsInfoService) CreateMallGoodsInfo(mallGoodsInfo manage.MallGoodsInfo) (err error) {
	var goodsCategory manage.MallGoodsCategory

	if !errors.Is(global.GVA_DB.Where("category_id=? AND deleted_status=0", goodsCategory.CategoryId).First(&goodsCategory).Error, gorm.ErrRecordNotFound) &&
		goodsCategory.CategoryLevel != enum.LevelThree.Code() {
		return errors.New("分类数据异常")
	}
	if !errors.Is(global.GVA_DB.Where("goods_name=? AND goods_category_id=?", mallGoodsInfo.GoodsName, mallGoodsInfo.GoodsCategoryId).First(&mallGoodsInfo).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的商品信息")
	}
	err = global.GVA_DB.Create(&mallGoodsInfo).Error
	return err
}

// DeleteMallGoodsInfo 删除MallGoodsInfo记录
func (m *MallGoodsInfoService) DeleteMallGoodsInfo(mallGoodsInfo manage.MallGoodsInfo) (err error) {
	err = global.GVA_DB.Delete(&mallGoodsInfo).Error
	return err
}

// ChangeMallGoodsInfoByIds 上下架
func (m *MallGoodsInfoService) ChangeMallGoodsInfoByIds(ids request.IdsReq, sellStatus string) (err error) {
	intSellStatus, _ := strconv.Atoi(sellStatus)
	err = global.GVA_DB.Model(&manage.MallGoodsInfo{}).Where("goods_id in ?", ids.Ids).Update("goods_sell_status", intSellStatus).Error
	return err
}

// UpdateMallGoodsInfo 更新MallGoodsInfo记录
func (m *MallGoodsInfoService) UpdateMallGoodsInfo(req manageReq.GoodsInfoUpdateParam) (err error) {
	goodsId, _ := strconv.Atoi(req.GoodsId)
	originalPrice, _ := strconv.Atoi(req.OriginalPrice)
	stockNum, _ := strconv.Atoi(req.StockNum)
	goodsInfo := manage.MallGoodsInfo{
		GoodsId:            goodsId,
		GoodsName:          req.GoodsName,
		GoodsIntro:         req.GoodsIntro,
		GoodsCategoryId:    req.GoodsCategoryId,
		GoodsCoverImg:      req.GoodsCoverImg,
		GoodsDetailContent: req.GoodsDetailContent,
		OriginalPrice:      originalPrice,
		SellingPrice:       req.SellingPrice,
		StockNum:           stockNum,
		Tag:                req.Tag,
		GoodsSellStatus:    req.GoodsSellStatus,
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}
	err = global.GVA_DB.Where("goods_id=?", goodsInfo.GoodsId).Updates(&goodsInfo).Error
	return err
}

// GetMallGoodsInfo 根据id获取MallGoodsInfo记录
func (m *MallGoodsInfoService) GetMallGoodsInfo(id int) (err error, mallGoodsInfo manage.MallGoodsInfo) {
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	return
}

// GetMallGoodsInfoInfoList 分页获取MallGoodsInfo记录
func (m *MallGoodsInfoService) GetMallGoodsInfoInfoList(info manageReq.MallGoodsInfoSearch, goodsName string, goodsSellStatus string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
	var mallGoodsInfos []manage.MallGoodsInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if goodsName != "" {
		db.Where("goods_name =?", goodsName)
	}
	if goodsSellStatus != "" {
		db.Where("goods_sell_status =?", goodsSellStatus)
	}
	err = db.Limit(limit).Offset(offset).Order("goods_id desc").Find(&mallGoodsInfos).Error
	return err, mallGoodsInfos, total
}