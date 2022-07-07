package initialize

import (
	_ "github.com/liujianjiang/goadmin/server/source/example"
	_ "github.com/liujianjiang/goadmin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
