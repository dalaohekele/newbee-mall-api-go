package manage

import (
	"errors"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
)

type MallUserService struct {
}

// CreateMallUser 创建MallUser记录
func (m *MallUserService) CreateMallUser(mallUser manage.MallUser) (err error) {
	err = global.GVA_DB.Create(&mallUser).Error
	return err
}

// DeleteMallUser 删除MallUser记录
func (m *MallUserService) DeleteMallUser(mallUser manage.MallUser) (err error) {
	err = global.GVA_DB.Delete(&mallUser).Error
	return err
}

// DeleteMallUserByIds 批量删除MallUser记录
func (m *MallUserService) DeleteMallUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]manage.MallUser{}, "id in ?", ids.Ids).Error
	return err
}

// LockUser 修改用户状态
func (m *MallUserService) LockUser(ids request.IdsReq, lockStatus int) (err error) {
	if lockStatus != 0 && lockStatus != 1 {
		return errors.New("操作非法！")
	}
	err = global.GVA_DB.Where("id in ?", ids).UpdateColumns(manage.MallUser{LockedFlag: lockStatus}).Error
	return err
}

// GetMallUser 根据id获取MallUser记录
func (m *MallUserService) GetMallUser(id uint) (err error, mallUser manage.MallUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&mallUser).Error
	return
}

// GetMallUserInfoList 分页获取商城注册用户列表
func (m *MallUserService) GetMallUserInfoList(info manageReq.MallUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallUser{})
	var mallUsers []manage.MallUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.LockedFlag == 1 || info.LockedFlag == 0 {
		db.Where("locked_flag=?", info.LockedFlag)
	}
	err = db.Limit(limit).Offset(offset).Find(&mallUsers).Error
	return err, mallUsers, total
}
