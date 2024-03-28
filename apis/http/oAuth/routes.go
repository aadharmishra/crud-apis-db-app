package oauth

import (
	"crud-apis-db-app/shared"

	"github.com/gin-gonic/gin"
)

func NewOAuthRouter(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewOAuthService(deps)
	routerApi := router.Group("/google")
	{
		routerApi.GET("/login", service.Login)
		routerApi.GET("/callback", service.Callback)
	}
}
