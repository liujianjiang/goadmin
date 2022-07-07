package service

import (
	"github.com/liujianjiang/goadmin/server/service/example"
	"github.com/liujianjiang/goadmin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
