package router

import "main.go/router/manage"

type RouterGroup struct {
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
