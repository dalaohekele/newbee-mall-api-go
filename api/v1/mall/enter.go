package mall

import "main.go/service"

type MallGroup struct {
	MallIndexApi
	MallGoodsInfoApi
	MallGoodsCategoryApi
	MallUserApi
}

var mallCarouselService = service.ServiceGroupApp.MallServiceGroup.MallCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.MallServiceGroup.MallGoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.MallServiceGroup.MallIndexInfoService
var mallGoodsCategoryService = service.ServiceGroupApp.MallServiceGroup.MallGoodsCategoryService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService
