package mall

import "main.go/service"

type MallGroup struct {
	MallIndexApi
}

var mallCarouselService = service.ServiceGroupApp.ManageServiceGroup.ManageCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.ManageServiceGroup.ManageIndexConfigService
