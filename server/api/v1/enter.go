package v1

import (
	"github.com/liujianjiang/goadmin/server/api/v1/example"
	"github.com/liujianjiang/goadmin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
