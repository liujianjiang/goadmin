package request

import (
	"github.com/liujianjiang/goadmin/server/model/{{.Package}}"
	"github.com/liujianjiang/goadmin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
