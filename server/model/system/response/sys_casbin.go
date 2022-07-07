package response

import (
	"github.com/liujianjiang/goadmin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
