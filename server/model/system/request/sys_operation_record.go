package request

import (
	"github.com/liujianjiang/goadmin/server/model/common/request"
	"github.com/liujianjiang/goadmin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
