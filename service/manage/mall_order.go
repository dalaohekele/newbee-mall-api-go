package manage

import (
	"errors"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"time"
)

type MallOrderService struct {
}

// CheckDone 修改订单状态为配货成功
func (mallOrderService *MallOrderService) CheckDone(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != 1 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err := global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 2, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功无法执行出库操作")
		}
	}
	return
}

// CheckOut 出库
func (mallOrderService *MallOrderService) CheckOut(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != 1 && order.OrderStatus != 2 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err := global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 3, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功或配货完成无法执行出库操作")
		}
	}
	return
}

// CloseOrder 商家关闭订单
func (mallOrderService *MallOrderService) CloseOrder(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus == 4 || order.OrderStatus < 0 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err := global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: -3, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单不能执行关闭操作")
		}
	}
	return
}

// GetMallOrder 根据id获取MallOrder记录
func (mallOrderService *MallOrderService) GetMallOrder(id int) (err error, mallOrder manage.MallOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&mallOrder).Error
	return
}

// GetMallOrderInfoList 分页获取MallOrder记录
func (mallOrderService *MallOrderService) GetMallOrderInfoList(info manageReq.MallOrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallOrder{})
	if info.OrderNo != "" {
		db.Where("order_no", info.OrderNo)
	}
	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
	if utils.NumsInList(info.OrderStatus, []int{-3, -2, -1, 0, 1, 2, 3, 4}) {
		db.Where("order_status", info.OrderStatus)
	}
	var mallOrders []manage.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&mallOrders).Error
	return err, mallOrders, total
}
