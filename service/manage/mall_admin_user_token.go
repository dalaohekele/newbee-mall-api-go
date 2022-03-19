package manage

import (
	"main.go/global"
	manageReq "main.go/model/common/request"
	"main.go/model/manage"
)

type MallAdminUserTokenService struct {
}

func (mallAdminUserTokenService *MallAdminUserTokenService) ExistAdminToken(token string) (err error, mallAdminUserToken manage.MallAdminUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallAdminUserToken).Error
	return
}

func (mallAdminUserTokenService *MallAdminUserTokenService) DeleteMallAdminUserTokenByIds(ids manageReq.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]manage.MallAdminUserToken{}, "id in ?", ids.Ids).Error
	return err
}
