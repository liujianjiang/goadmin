package request

import (
	"github.com/liujianjiang/goadmin/server/model/common/request"
	"github.com/liujianjiang/goadmin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
