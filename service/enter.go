package service

import (
	"main.go/service/example"
	"main.go/service/manage"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
