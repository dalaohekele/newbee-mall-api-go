package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	"time"
)

type MallUserAddressService struct {
}

func (m *MallUserAddressService) GetMyAddress(token string) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	err = global.GVA_DB.Where("user_id=?", userToken.UserId).First(&userAddress).Error
	return
}

func (m *MallUserAddressService) SaveUserAddress(token string, req mallReq.AddAddressParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	// 是否新增了默认地址，将之前的默认地址设置为非默认
	if req.DefaultFlag == 1 {
		var defaultAddress mall.MallUserAddress
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultAddress)
		if defaultAddress != (mall.MallUserAddress{}) {
			defaultAddress.DefaultFlag = 0
			defaultAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultAddress).Error
			if err != nil {
				return
			}
		}
	}
	err = global.GVA_DB.Create(&req).Error
	return
}

func (m *MallUserAddressService) UpdateUserAddress(token string, req mallReq.UpdateAddressParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", req.AddressId).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	err = copier.Copy(&userAddress, &req)
	if err != nil {
		return
	}
	userAddress.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&userAddress).Error
	return
}

func (m *MallUserAddressService) GetMallUserAddressById(token string, id int) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址"), userAddress
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！"), userAddress
	}
	return
}

func (m *MallUserAddressService) GetMallUserDefaultAddress(token string) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("不存在默认地址失败"), userAddress
	}
	return
}

func (m *MallUserAddressService) DeleteUserAddress(token string, id int) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}
