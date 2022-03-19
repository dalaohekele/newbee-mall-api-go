package v1

import "main.go/api/v1/manage"

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
}

var ApiGroupApp = new(ApiGroup)
