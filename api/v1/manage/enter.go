package manage

import "main.go/service"

type ManageGroup struct {
	MallAdminUserApi
	GoodsCategoryApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.MallAdminUserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.MallAdminUserTokenService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.MallUserService
var goodsCategoryService = service.ServiceGroupApp.ManageServiceGroup.GoodsCategoryService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
