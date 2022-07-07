package router

import (
	"github.com/liujianjiang/goadmin/server/router/example"
	"github.com/liujianjiang/goadmin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
