package mall

import "main.go/service"

type MallGroup struct {
	MallIndexApi
}

var mallCarouselService = service.ServiceGroupApp.ManageServiceGroup.MallCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.ManageServiceGroup.MallGoodsInfoService
