package manage

import (
	"main.go/global"
	manageReq "main.go/model/common/request"
	"main.go/model/manage"
)

type ManageAdminUserTokenService struct {
}

func (m *ManageAdminUserTokenService) ExistAdminToken(token string) (err error, mallAdminUserToken manage.MallAdminUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallAdminUserToken).Error
	return
}

func (m *ManageAdminUserTokenService) DeleteMallAdminUserTokenByIds(ids manageReq.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]manage.MallAdminUserToken{}, "id in ?", ids.Ids).Error
	return err
}
