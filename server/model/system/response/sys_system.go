package response

import "github.com/liujianjiang/goadmin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
