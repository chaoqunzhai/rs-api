package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/site/apis"
)

func init() {

	routerNoCheckRole = append(routerNoCheckRole, PublicNoAuthRouter)
}

func PublicNoAuthRouter(v1 *gin.RouterGroup) {
	api := apis.Config{}
	r := v1.Group("/site")
	{
		r.GET("loginConfig", api.Config)
		r.GET("captcha", api.Captcha)
	}
}
