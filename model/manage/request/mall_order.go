package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallOrderSearch struct {
	manage.MallOrder
	request.PageInfo
}
