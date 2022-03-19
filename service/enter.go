package service

import "main.go/service/manage"

type ServiceGroup struct {
	ManageServiceGroup manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
