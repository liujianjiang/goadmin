package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liujianjiang/goadmin/server/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.GET("send_email", baseApi.SendEmail)
	}
	return baseRouter
}
