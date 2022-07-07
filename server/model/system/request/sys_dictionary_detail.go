package request

import (
	"github.com/liujianjiang/goadmin/server/model/common/request"
	"github.com/liujianjiang/goadmin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
