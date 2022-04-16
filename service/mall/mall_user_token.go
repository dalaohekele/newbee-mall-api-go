package mall

import (
	"main.go/global"
	manageReq "main.go/model/common/request"
	"main.go/model/mall"
)

type MallUserTokenService struct {
}

func (m *MallUserTokenService) ExistUserToken(token string) (err error, mallUserToken mall.MallUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *MallUserTokenService) DeleteMallUserTokenByIds(ids manageReq.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mall.MallUserToken{}, "user_id in ?", ids.Ids).Error
	return err
}
